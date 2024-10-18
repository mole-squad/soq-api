package timewindows

import (
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TimeWindowRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type TimeWindowRepoResult struct {
	fx.Out

	TimeWindowRepo interfaces.TimeWindowRepo
}

type TimeWindowRepo struct {
	*generics.Repository[*models.TimeWindow]
}

func NewTimeWindowRepo(params TimeWindowRepoParams) (TimeWindowRepoResult, error) {
	embeddedRepo := generics.NewRepository[*models.TimeWindow](
		params.DBService,
		params.LoggerService,
		generics.WithTableName[*models.TimeWindow]("time_windows"),
	).(*generics.Repository[*models.TimeWindow])

	repo := &TimeWindowRepo{
		Repository: embeddedRepo,
	}

	return TimeWindowRepoResult{TimeWindowRepo: repo}, nil
}
