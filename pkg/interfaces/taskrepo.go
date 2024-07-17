package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type TaskRepo interface {
	CreateOne(ctx context.Context, task *models.Task) error
	UpdateOne(ctx context.Context, task *models.Task) error
	DeleteOne(ctx context.Context, id uint) error
	FindOneByUser(ctx context.Context, userID uint, query string, args ...interface{}) (*models.Task, error)
	FindManyByUser(ctx context.Context, userID uint, query string, args ...interface{}) ([]*models.Task, error)
}
