package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type Repo[M models.Model] interface {
	FindOneByUser(ctx context.Context, userID uint, query string, args ...interface{}) (M, error)
	FindManyByUser(ctx context.Context, userID uint, query string, args ...interface{}) ([]M, error)
	CreateOne(ctx context.Context, item M) error
	UpdateOne(ctx context.Context, item M) error
	DeleteOne(ctx context.Context, itemID uint) error
}
