package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type TaskRepo interface {
	CreateOne(ctx context.Context, task *models.Task) error
	UpdateOne(ctx context.Context, task *models.Task) error
	FindManyByUser(ctx context.Context, userID uint) ([]models.Task, error)
}
