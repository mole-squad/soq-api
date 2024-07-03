package tasks

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/interfaces"
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
	taskRepo interfaces.TaskRepo
}

func NewTaskService(params TaskServiceParams) (TaskServiceResult, error) {
	srv := &TaskService{taskRepo: params.TaskRepo}
	return TaskServiceResult{TaskService: srv}, nil
}

func (srv *TaskService) ListUserTasks(ctx context.Context, user interfaces.User) ([]interfaces.Task, error) {
	tasks, err := srv.taskRepo.FindManyTasksByUser(ctx, user.ID())
	if err != nil {
		return []interfaces.Task{}, fmt.Errorf("failed to list user tasks: %w", err)
	}

	return tasks, nil
}
