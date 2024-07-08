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
	&models.TimeWindow{},
	&models.FocusArea{},
	&models.Task{},
	&models.Quota{},
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

	err = srv.Migrate(context.Background())
	if err != nil {
		return fmt.Errorf("migrate failed: %w", err)
	}

	return nil
}

func (srv *DBService) CreateOne(ctx context.Context, record interface{}) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	createResult := sesh.Create(record)
	if createResult.Error != nil {
		return fmt.Errorf("create one failed: %w", createResult.Error)
	}

	return nil
}

func (srv *DBService) UpdateOne(ctx context.Context, record interface{}) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	updateResult := sesh.Model(record).Updates(record)
	if updateResult.Error != nil {
		return fmt.Errorf("update one failed: %w", updateResult.Error)
	}

	return nil
}

func (srv *DBService) DeleteOne(ctx context.Context, record interface{}) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	deleteResult := sesh.Delete(record)
	if deleteResult.Error != nil {
		return fmt.Errorf("delete one failed: %w", deleteResult.Error)
	}

	return nil
}

func (srv *DBService) FindMany(
	ctx context.Context,
	result interface{},
	joins []string,
	preloads []string,
	query interface{},
	args ...interface{},
) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	for _, join := range joins {
		sesh = sesh.Joins(join)
	}

	for _, preload := range preloads {
		sesh = sesh.Preload(preload)
	}

	queryResult := sesh.Where(query, args).Find(result)
	if queryResult.Error != nil {
		return fmt.Errorf("find many failed: %w", queryResult.Error)
	}

	return nil
}

func (srv *DBService) Migrate(ctx context.Context) error {
	for _, model := range allModels {
		if err := srv.db.AutoMigrate(model); err != nil {
			return fmt.Errorf("migrate failed for model %v: %w", model, err)
		}
	}

	return nil
}

func (srv *DBService) DropAll(ctx context.Context) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	for _, model := range allModels {
		err := sesh.Migrator().DropTable(model)
		if err != nil {
			return fmt.Errorf("drop all failed: %w", err)
		}
	}

	return nil
}

func (srv *DBService) GetSession(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	timeoutCtx, cancel := context.WithTimeout(ctx, QueryTimeout)

	return srv.db.Session(&gorm.Session{
		Context: timeoutCtx,
	}), cancel
}
