package db

import (
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/burkel24/task-app/pkg/users"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBService struct {
}

var _ interfaces.DBService = DBService{}

func NewDBService() interfaces.DBService {
	return DBService{}
}

var Module = fx.Module(
	"DB",
	fx.Provide(NewDBService),
)

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=pass dbname=task port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&tasks.Task{})

	return db, err
}
