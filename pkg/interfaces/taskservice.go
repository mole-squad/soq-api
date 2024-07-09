package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type TaskService interface {
	CreateUserTask(ctx context.Context, user *models.User, task *models.Task) (models.Task, error)
	UpdateUserTask(ctx context.Context, task *models.Task) (models.Task, error)
	DeleteUserTask(ctx context.Context, id uint) error
	ListOpenUserTasks(ctx context.Context, userID uint) ([]models.Task, error)
}
