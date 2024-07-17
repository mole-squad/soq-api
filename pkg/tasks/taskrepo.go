package tasks

import (
	"github.com/mole-squad/soq-api/pkg/db"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

var (
	joins = []string{"FocusArea"}
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
	db.Repo[*models.Task]

	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewTaskRepo(params TaskRepoParams) (TaskRepoResult, error) {
	embeddedRepo := db.NewRepo[*models.Task](
		params.DBService,
		params.LoggerService,
		db.WithTableName[*models.Task]("tasks"),
		db.WithJoinTables[*models.Task]("FocusArea"),
	)

	repo := &TaskRepo{
		Repo:      *embeddedRepo,
		dbService: params.DBService,
		logger:    params.LoggerService,
	}

	return TaskRepoResult{TaskRepo: repo}, nil
}
