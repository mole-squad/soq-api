package tasks

import (
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
	srv := TaskService{taskRepo: params.TaskRepo}
	return TaskServiceResult{TaskService: srv}, nil
}
