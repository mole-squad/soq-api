package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type TaskService interface {
	ResourceService[*models.Task]

	ResolveTask(ctx context.Context, taskID uint) (*models.Task, error)
	ListOpenUserTasksForFocusArea(ctx context.Context, userID uint, focusAreaID uint) ([]*models.Task, error)
}
