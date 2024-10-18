package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TimeWindowControllerParams struct {
	fx.In

	AuthService       interfaces.AuthService
	LoggerService     interfaces.LoggerService
	TimeWindowService interfaces.TimeWindowService
	Router            *chi.Mux
}

type TimeWindowControllerResult struct {
	fx.Out

	TimeWindowController TimeWindowController
}

type TimeWindowController struct {
	interfaces.Controller[*models.TimeWindow]
}

func NewTimeWindowController(params TimeWindowControllerParams) (TimeWindowControllerResult, error) {
	ctrl := TimeWindowController{}

	ctrl.Controller = generics.NewController[*models.TimeWindow](
		params.TimeWindowService,
		params.LoggerService,
		params.AuthService,
		models.NewTimeWindowFromCreateRequest,
		models.NewTimeWindowFromUpdateRequest,
		generics.WithContextKey[*models.TimeWindow](timeWindowContextKey),
	).(*generics.Controller[*models.TimeWindow])

	params.Router.Mount("/timewindows", ctrl.Controller.GetRouter())

	return TimeWindowControllerResult{TimeWindowController: ctrl}, nil
}
