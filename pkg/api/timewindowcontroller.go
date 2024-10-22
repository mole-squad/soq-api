package api

import (
	"github.com/burkel24/go-mochi"

	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type TimeWindowControllerParams struct {
	fx.In

	AuthService       mochi.AuthService
	LoggerService     mochi.LoggerService
	TimeWindowService interfaces.TimeWindowService
	Router            *chi.Mux
}

type TimeWindowControllerResult struct {
	fx.Out

	TimeWindowController TimeWindowController
}

type TimeWindowController struct {
	mochi.Controller[*models.TimeWindow]
}

func NewTimeWindowController(params TimeWindowControllerParams) (TimeWindowControllerResult, error) {
	ctrl := TimeWindowController{}

	ctrl.Controller = mochi.NewController(
		params.TimeWindowService,
		params.LoggerService,
		params.AuthService,
		models.NewTimeWindowFromCreateRequest,
		models.NewTimeWindowFromUpdateRequest,
		mochi.WithContextKey[*models.TimeWindow](timeWindowContextKey),
	)

	params.Router.Mount("/timewindows", ctrl.Controller.GetRouter())

	return TimeWindowControllerResult{TimeWindowController: ctrl}, nil
}
