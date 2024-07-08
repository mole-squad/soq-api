package tasks

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TaskRepoParams struct {
	fx.In

	DBService interfaces.DBService
}

type TaskRepoResult struct {
	fx.Out

	TaskRepo interfaces.TaskRepo
}

type TaskRepo struct {
	dbService interfaces.DBService
}

func NewTaskRepo(params TaskRepoParams) (TaskRepoResult, error) {
	repo := &TaskRepo{dbService: params.DBService}
	return TaskRepoResult{TaskRepo: repo}, nil
}

func (repo *TaskRepo) CreateOne(ctx context.Context, task *models.Task) error {
	slog.Info("Creating one task", "task", task)

	err := repo.dbService.CreateOne(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to create one task: %w", err)
	}

	return nil
}

func (repo *TaskRepo) UpdateOne(ctx context.Context, task *models.Task) error {
	slog.Info("Updating one task", "task", task)

	err := repo.dbService.UpdateOne(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to update one task: %w", err)
	}

	return nil
}

func (repo *TaskRepo) DeleteOne(ctx context.Context, id uint) error {
	slog.Info("Deleting one task", "id", id)

	task := &models.Task{Model: gorm.Model{ID: id}}

	err := repo.dbService.DeleteOne(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to delete one task: %w", err)
	}

	return nil
}

func (repo *TaskRepo) FindManyByUser(ctx context.Context, userID uint) ([]models.Task, error) {
	var tasks []models.Task

	err := repo.dbService.FindMany(ctx, &tasks, []string{"FocusArea"}, "tasks.user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find many taks by user: %w", err)
	}

	return tasks, nil
}
