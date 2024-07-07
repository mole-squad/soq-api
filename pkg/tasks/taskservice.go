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
	task.UserID = user.ID

	err := srv.taskRepo.CreateOne(ctx, task)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to create user task: %w", err)
	}

	return *task, nil
}

func (srv *TaskService) UpdateUserTask(
	ctx context.Context,
	task *models.Task,
) (models.Task, error) {
	err := srv.taskRepo.UpdateOne(ctx, task)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to update user task: %w", err)
	}

	return *task, nil
}

func (srv *TaskService) DeleteUserTask(ctx context.Context, id uint) error {
	err := srv.taskRepo.DeleteOne(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user task: %w", err)
	}

	return nil
}

func (srv *TaskService) ListUserTasks(ctx context.Context, user *models.User) ([]models.Task, error) {
	tasks, err := srv.taskRepo.FindManyByUser(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to list user tasks: %w", err)
	}

	return tasks, nil
}
