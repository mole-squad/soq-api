package timewindows

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TimeWindowRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService mochi.LoggerService
}

type TimeWindowRepoResult struct {
	fx.Out

	TimeWindowRepo interfaces.TimeWindowRepo
}

type TimeWindowRepo struct {
	mochi.Repository[*models.TimeWindow]
}

func NewTimeWindowRepo(params TimeWindowRepoParams) (TimeWindowRepoResult, error) {
	embeddedRepo := mochi.NewRepository(
		params.DBService,
		params.LoggerService,
		mochi.WithTableName[*models.TimeWindow]("time_windows"),
	)

	repo := &TimeWindowRepo{
		Repository: embeddedRepo,
	}

	return TimeWindowRepoResult{TimeWindowRepo: repo}, nil
}
