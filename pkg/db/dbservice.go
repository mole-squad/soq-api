package db

import (
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/burkel24/task-app/pkg/users"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var models = []interface{}{
	&users.User{},
	&tasks.Task{},
}

type DBServiceParams struct {
	fx.In
}

type DbServiceResult struct {
	fx.Out

	DBService interfaces.DBService
}

type DBService struct {
	db *gorm.DB
}

func NewDBService() (DbServiceResult, error) {
	srv := DBService{}
	srv.Init()

	return DbServiceResult{DBService: srv}, nil
}

func (srv *DBService) Init() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=pass dbname=task port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	srv.db = db

	for _, model := range models {
		srv.db.AutoMigrate(model)
	}

	return err
}
