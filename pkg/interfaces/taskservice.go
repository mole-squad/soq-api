package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type TaskService interface {
	CreateOne(ctx context.Context, userID uint, task *models.Task) (*models.Task, error)
	GetOne(ctx context.Context, userID, taskID uint) (*models.Task, error)
	UpdateOne(ctx context.Context, userID uint, taskID uint, task *models.Task) (*models.Task, error)
	ResolveUserTask(ctx context.Context, userID uint, taskID uint) (models.Task, error)
	DeleteOne(ctx context.Context, userID uint, taskID uint) error
	List(ctx context.Context, userID uint) ([]*models.Task, error)
	ListOpenUserTasksForFocusArea(ctx context.Context, userID uint, focusAreaID uint) ([]*models.Task, error)
}
