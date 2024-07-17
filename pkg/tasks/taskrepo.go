package tasks

import (
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
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
	*generics.ResourceRepository[*models.Task]

	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewTaskRepo(params TaskRepoParams) (TaskRepoResult, error) {
	embeddedRepo := generics.NewResourceRepository[*models.Task](
		params.DBService,
		params.LoggerService,
		generics.WithTableName[*models.Task]("tasks"),
		generics.WithJoinTables[*models.Task]("FocusArea"),
	).(*generics.ResourceRepository[*models.Task])

	repo := &TaskRepo{
		ResourceRepository: embeddedRepo,
		dbService:          params.DBService,
		logger:             params.LoggerService,
	}

	return TaskRepoResult{TaskRepo: repo}, nil
}
