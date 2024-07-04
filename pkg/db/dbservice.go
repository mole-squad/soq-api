package db

import (
	"context"
	"fmt"
	"time"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var allModels = []interface{}{
	&models.User{},
	&models.Task{},
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

	for _, model := range allModels {
		srv.db.AutoMigrate(model)
	}

	return err
}

func (srv *DBService) CreateOne(ctx context.Context, record interface{}) error {
	sesh, cancel := srv.buildSession(ctx)
	defer cancel()

	createResult := sesh.Create(record)
	if createResult.Error != nil {
		return fmt.Errorf("create one failed: %w", createResult.Error)
	}

	return nil
}

func (srv *DBService) UpdateOne(ctx context.Context, record interface{}) error {
	sesh, cancel := srv.buildSession(ctx)
	defer cancel()

	updateResult := sesh.Model(record).Updates(record)
	if updateResult.Error != nil {
		return fmt.Errorf("update one failed: %w", updateResult.Error)
	}

	return nil
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
