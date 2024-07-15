package tasks

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
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

func (srv *TaskService) ResolveUserTask(
	ctx context.Context,
	userID uint,
	taskID uint,
) (models.Task, error) {
	task := models.Task{
		Model:  gorm.Model{ID: taskID},
		UserID: userID,
		Status: models.TaskStatusClosed,
	}

	err := srv.taskRepo.UpdateOne(ctx, &task)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to resolve user task: %w", err)
	}

	return task, nil
}

func (srv *TaskService) DeleteUserTask(ctx context.Context, id uint) error {
	err := srv.taskRepo.DeleteOne(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user task: %w", err)
	}

	return nil
}

func (srv *TaskService) ListOpenUserTasks(ctx context.Context, userID uint) ([]models.Task, error) {
	tasks, err := srv.taskRepo.FindManyByUser(ctx, userID, "status = ?", models.TaskStatusOpen)
	if err != nil {
		return nil, fmt.Errorf("failed to list user tasks: %w", err)
	}

	return tasks, nil
}

func (srv *TaskService) ListOpenUserTasksForFocusArea(ctx context.Context, userID uint, focusAreaID uint) ([]models.Task, error) {
	tasks, err := srv.taskRepo.FindManyByUser(ctx, userID, "status = ? AND focus_area_id = ?", models.TaskStatusOpen, focusAreaID)
	if err != nil {
		return nil, fmt.Errorf("failed to list user tasks: %w", err)
	}

	return tasks, nil
}
