package tasks

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/interfaces"
	"go.uber.org/fx"
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

func (repo *TaskRepo) FindManyTasksByUser(ctx context.Context, userID uint) ([]interfaces.Task, error) {
	var tasks []Task

	err := repo.dbService.FindMany(ctx, &tasks, "user_id = ?", userID)
	if err != nil {
		return []interfaces.Task{}, fmt.Errorf("failed to find many taks by user: %w", err)
	}

	taskResult := make([]interfaces.Task, len(tasks))
	for idx, task := range tasks {
		taskResult[idx] = &task
	}

	return taskResult, nil
}
