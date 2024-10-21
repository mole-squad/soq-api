package api

import (
	"github.com/burkel24/go-mochi"

	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type FocusAreaControllerParams struct {
	fx.In

	AuthService      mochi.AuthService
	FocusAreaService interfaces.FocusAreaService
	LoggerService    mochi.LoggerService
	Router           *chi.Mux
}

type FocusAreaControllerResult struct {
	fx.Out

	FocusAreaController FocusAreaController
}

type FocusAreaController struct {
	mochi.Controller[*models.FocusArea]
}

func NewFocusAreaController(params FocusAreaControllerParams) (FocusAreaControllerResult, error) {
	ctrl := FocusAreaController{}

	ctrl.Controller = mochi.NewController(
		params.FocusAreaService,
		params.LoggerService,
		params.AuthService,
		models.NewFocusAreaFromCreateRequest,
		models.NewFocusAreaFromUpdateRequest,
		mochi.WithContextKey[*models.FocusArea](focusAreaContextKey),
	)

	params.Router.Mount("/focusareas", ctrl.Controller.GetRouter())

	return FocusAreaControllerResult{FocusAreaController: ctrl}, nil
}
