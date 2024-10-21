package api

import (
	"github.com/burkel24/go-mochi"

	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type DeviceControllerParams struct {
	fx.In

	AuthService   mochi.AuthService
	DeviceService interfaces.DeviceService
	LoggerService mochi.LoggerService
	Router        *chi.Mux
}

type DeviceControllerResult struct {
	fx.Out

	DeviceController DeviceController
}

type DeviceController struct {
	mochi.Controller[*models.Device]
}

func NewDeviceController(params DeviceControllerParams) (DeviceControllerResult, error) {
	ctrl := DeviceController{}

	ctrl.Controller = mochi.NewController(
		params.DeviceService,
		params.LoggerService,
		params.AuthService,
		models.NewDeviceFromCreateRequest,
		models.NewDeviceFromUpdateRequest,
		mochi.WithContextKey[*models.Device](deviceContextKey),
	)

	params.Router.Mount("/devices", ctrl.Controller.GetRouter())

	return DeviceControllerResult{DeviceController: ctrl}, nil
}
