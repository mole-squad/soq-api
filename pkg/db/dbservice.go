package db

import (
	"context"
	"fmt"
	"time"

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

const (
	QueryTimeout = time.Second
)

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
	srv := &DBService{}
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

func (srv *DBService) FindMany(ctx context.Context, result interface{}, query interface{}, args ...interface{}) error {
	sesh, cancel := srv.buildSession(ctx)
	defer cancel()

	queryResult := sesh.Where(query, args).Find(result)
	if queryResult.Error != nil {
		return fmt.Errorf("find many failed: %w", queryResult.Error)
	}

	return nil
}

func (srv *DBService) buildSession(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	timeoutCtx, cancel := context.WithTimeout(ctx, QueryTimeout)

	return srv.db.Session(&gorm.Session{
		Context: timeoutCtx,
	}), cancel
}
