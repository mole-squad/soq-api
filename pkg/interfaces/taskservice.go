package interfaces

import (
	"context"

	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/models"
)

type TaskService interface {
	mochi.Service[*models.Task]

	ResolveTask(ctx context.Context, taskID uint) (*models.Task, error)
	ListOpenUserTasksForFocusArea(ctx context.Context, userID uint, focusAreaID uint) ([]*models.Task, error)
}
