package timewindows

import (
	"github.com/mole-squad/soq-api/pkg/generics"
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
	*generics.Service[*models.TimeWindow]
}

func NewTimeWindowService(params TimeWindowServiceParams) (TimeWindowServiceResult, error) {
	embeddedSvc := generics.NewService[*models.TimeWindow](
		params.TimeWindowRepo,
	).(*generics.Service[*models.TimeWindow])

	srv := &TimeWindowService{
		Service: embeddedSvc,
	}

	return TimeWindowServiceResult{TimeWindowService: srv}, nil
}
