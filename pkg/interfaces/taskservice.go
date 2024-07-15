package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type TaskService interface {
	CreateUserTask(ctx context.Context, user *models.User, task *models.Task) (models.Task, error)
	UpdateUserTask(ctx context.Context, task *models.Task) (models.Task, error)
	ResolveUserTask(ctx context.Context, userID uint, taskID uint) (models.Task, error)
	DeleteUserTask(ctx context.Context, id uint) error
	ListOpenUserTasks(ctx context.Context, userID uint) ([]models.Task, error)
	ListOpenUserTasksForFocusArea(ctx context.Context, userID uint, focusAreaID uint) ([]models.Task, error)
}
