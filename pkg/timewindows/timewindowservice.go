package timewindows

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TimeWindowServiceParams struct {
	fx.In

	TimeWindowRepo interfaces.TimeWindowRepo
}

type TimeWindowServiceResult struct {
	fx.Out

	TimeWindowService interfaces.TimeWindowService
}

type TimeWindowService struct {
	mochi.Service[*models.TimeWindow]
}

func NewTimeWindowService(params TimeWindowServiceParams) (TimeWindowServiceResult, error) {
	embeddedSvc := mochi.NewService(
		params.TimeWindowRepo,
	)

	srv := &TimeWindowService{
		Service: embeddedSvc,
	}

	return TimeWindowServiceResult{TimeWindowService: srv}, nil
}
