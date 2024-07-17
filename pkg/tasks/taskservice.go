package tasks

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TaskServiceParams struct {
	fx.In

	TaskRepo interfaces.TaskRepo
}

type TaskServiceResult struct {
	fx.Out

	TaskService interfaces.TaskService
}

type TaskService struct {
	*generics.ResourceService[*models.Task]

	taskRepo interfaces.TaskRepo
}

func NewTaskService(params TaskServiceParams) (TaskServiceResult, error) {
	embeddedSvc := generics.NewResourceService[*models.Task](
		params.TaskRepo,
		generics.WithListQuery[*models.Task]("status = ?", models.TaskStatusOpen),
	).(*generics.ResourceService[*models.Task])

	srv := &TaskService{
		ResourceService: embeddedSvc,
		taskRepo:        params.TaskRepo,
	}

	return TaskServiceResult{TaskService: srv}, nil
}

func (srv *TaskService) ResolveTask(ctx context.Context, taskID uint) (*models.Task, error) {
	task := &models.Task{
		Status: models.TaskStatusClosed,
	}

	err := srv.taskRepo.UpdateOne(ctx, taskID, task)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve user task: %w", err)
	}

	return task, nil
}

func (srv *TaskService) ListOpenUserTasksForFocusArea(
	ctx context.Context,
	userID uint,
	focusAreaID uint,
) ([]*models.Task, error) {
	tasks, err := srv.taskRepo.FindManyByUser(ctx, userID, "status = ? AND focus_area_id = ?", models.TaskStatusOpen, focusAreaID)
	if err != nil {
		return nil, fmt.Errorf("failed to list user tasks: %w", err)
	}

	return tasks, nil
}
