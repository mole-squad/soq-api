package tasks

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq/pkg/interfaces"
	"github.com/mole-squad/soq/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TaskRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type TaskRepoResult struct {
	fx.Out

	TaskRepo interfaces.TaskRepo
}

type TaskRepo struct {
	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewTaskRepo(params TaskRepoParams) (TaskRepoResult, error) {
	repo := &TaskRepo{
		dbService: params.DBService,
		logger:    params.LoggerService,
	}

	return TaskRepoResult{TaskRepo: repo}, nil
}

func (repo *TaskRepo) CreateOne(ctx context.Context, task *models.Task) error {
	repo.logger.Info("Creating one task", "task", task)

	err := repo.dbService.CreateOne(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to create one task: %w", err)
	}

	return nil
}

func (repo *TaskRepo) UpdateOne(ctx context.Context, task *models.Task) error {
	repo.logger.Info("Updating one task", "task", task)

	err := repo.dbService.UpdateOne(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to update one task: %w", err)
	}

	return nil
}

func (repo *TaskRepo) DeleteOne(ctx context.Context, id uint) error {
	repo.logger.Info("Deleting one task", "id", id)

	task := &models.Task{Model: gorm.Model{ID: id}}

	err := repo.dbService.DeleteOne(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to delete one task: %w", err)
	}

	return nil
}

func (repo *TaskRepo) FindManyByUser(ctx context.Context, userID uint, query string, args ...interface{}) ([]models.Task, error) {
	var tasks []models.Task

	fullQuery := "tasks.user_id = ?"
	if query != "" {
		fullQuery = fmt.Sprintf("%s AND %s", fullQuery, query)
	}

	fullArgs := append([]interface{}{userID}, args...)

	err := repo.dbService.FindMany(ctx, &tasks, []string{"FocusArea"}, []string{}, fullQuery, fullArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to find many taks by user: %w", err)
	}

	return tasks, nil
}
