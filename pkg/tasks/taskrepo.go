package tasks

import (
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
	repo := TaskRepo{dbService: params.DBService}
	return TaskRepoResult{TaskRepo: repo}, nil
}
