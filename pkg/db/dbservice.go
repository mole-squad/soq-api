package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var allModels = []interface{}{
	&models.User{},
	&models.TimeWindow{},
	&models.FocusArea{},
	&models.Task{},
	&models.Quota{},
	&models.Agenda{},
	&models.AgendaItem{},
	&models.Device{},
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
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUrl,
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

func (srv *DBService) UpdateOne(ctx context.Context, recordID uint, record interface{}) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	updateResult := sesh.
		Model(record).
		Where("id = ?", recordID).
		Clauses(clause.Returning{}).
		Updates(record)

	if updateResult.Error != nil {
		return fmt.Errorf("update one failed: %w", updateResult.Error)
	}

	return nil
}

func (srv *DBService) DEPUpdateOne(ctx context.Context, record interface{}) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	updateResult := sesh.Model(record).Clauses(clause.Returning{}).Updates(record)
	if updateResult.Error != nil {
		return fmt.Errorf("update one failed: %w", updateResult.Error)
	}

	return nil
}

func (srv *DBService) DeleteOne(ctx context.Context, recordID uint, record interface{}) error {
	sesh, cancel := srv.GetSession(ctx)
	defer cancel()

	deleteResult := sesh.Delete(record, recordID)
	if deleteResult.Error != nil {
		return fmt.Errorf("delete one failed: %w", deleteResult.Error)
	}

	return nil
}

func (srv *DBService) FindOne(
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

	if query != nil {
		sesh = sesh.Where(query, args...)
	}

	queryResult := sesh.First(result)
	if queryResult.Error != nil {
		if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
			return ErrorNotFound
		}

		return fmt.Errorf("find one failed: %w", queryResult.Error)
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

	if query != nil {
		sesh = sesh.Where(query, args...)
	}

	queryResult := sesh.Find(result)
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
