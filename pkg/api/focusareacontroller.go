package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type FocusAreaControllerParams struct {
	fx.In

	AuthService      interfaces.AuthService
	FocusAreaService interfaces.FocusAreaService
	LoggerService    interfaces.LoggerService
	Router           *chi.Mux
}

type FocusAreaControllerResult struct {
	fx.Out

	FocusAreaController FocusAreaController
}

type FocusAreaController struct {
	interfaces.Controller[*models.FocusArea]
}

func NewFocusAreaController(params FocusAreaControllerParams) (FocusAreaControllerResult, error) {
	ctrl := FocusAreaController{}

	ctrl.Controller = generics.NewController[*models.FocusArea](
		params.FocusAreaService,
		params.LoggerService,
		params.AuthService,
		models.NewFocusAreaFromCreateRequest,
		models.NewFocusAreaFromUpdateRequest,
		generics.WithContextKey[*models.FocusArea](focusAreaContextKey),
	).(*generics.Controller[*models.FocusArea])

	params.Router.Mount("/focusareas", ctrl.Controller.GetRouter())

	return FocusAreaControllerResult{FocusAreaController: ctrl}, nil
}
