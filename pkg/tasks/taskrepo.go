package tasks

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TaskRepoParams struct {
	fx.In

	DBService     mochi.DBService
	LoggerService mochi.LoggerService
}

type TaskRepoResult struct {
	fx.Out

	TaskRepo interfaces.TaskRepo
}

type TaskRepo struct {
	mochi.Repository[*models.Task]

	dbService mochi.DBService
	logger    mochi.LoggerService
}

func NewTaskRepo(params TaskRepoParams) (TaskRepoResult, error) {
	embeddedRepo := mochi.NewRepository(
		params.DBService,
		params.LoggerService,
		mochi.WithTableName[*models.Task]("tasks"),
		mochi.WithJoinTables[*models.Task]("FocusArea"),
	)

	repo := &TaskRepo{
		Repository: embeddedRepo,
		dbService:  params.DBService,
		logger:     params.LoggerService,
	}

	return TaskRepoResult{TaskRepo: repo}, nil
}
