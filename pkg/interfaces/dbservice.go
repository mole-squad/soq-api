package interfaces

import (
	"context"

	"gorm.io/gorm"
)

type DBService interface {
	CreateOne(ctx context.Context, record interface{}) error
	UpdateOne(ctx context.Context, record interface{}) error
	DeleteOne(ctx context.Context, record interface{}) error
	FindOne(
		ctx context.Context,
		result interface{},
		joins []string,
		preloads []string,
		query interface{},
		args ...interface{},
	) error
	FindMany(
		ctx context.Context,
		result interface{},
		joins []string,
		preloads []string,
		query interface{},
		args ...interface{},
	) error

	GetSession(ctx context.Context) (*gorm.DB, context.CancelFunc)
	Migrate(ctx context.Context) error
	DropAll(ctx context.Context) error
}
