package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type ResourceRepository[M models.Model] interface {
	FindOne(ctx context.Context, query string, args ...interface{}) (M, error)
	FindOneByID(ctx context.Context, itemID uint, query string, args ...interface{}) (M, error)
	FindOneByUser(ctx context.Context, userID uint, query string, args ...interface{}) (M, error)
	FindManyByUser(ctx context.Context, userID uint, query string, args ...interface{}) ([]M, error)
	CreateOne(ctx context.Context, item M) error
	UpdateOne(ctx context.Context, itemID uint, item M) error
	DeleteOne(ctx context.Context, itemID uint) error
}
