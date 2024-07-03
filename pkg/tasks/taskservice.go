package tasks

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
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

func (srv *TaskService) CreateUserTask(
	ctx context.Context,
	user *models.User,
	task *models.Task,
) (models.Task, error) {
	newTask := models.Task{
		UserID:  user.ID,
		Summary: task.Summary,
	}

	err := srv.taskRepo.CreateOne(ctx, &newTask)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to create user task: %w", err)
	}

	return newTask, nil
}

func (srv *TaskService) ListUserTasks(ctx context.Context, user *models.User) ([]models.Task, error) {
	tasks, err := srv.taskRepo.FindManyByUser(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to list user tasks: %w", err)
	}

	return tasks, nil
}
