package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type DeviceControllerParams struct {
	fx.In

	AuthService   interfaces.AuthService
	DeviceService interfaces.DeviceService
	LoggerService interfaces.LoggerService
	Router        *chi.Mux
}

type DeviceControllerResult struct {
	fx.Out

	DeviceController DeviceController
}

type DeviceController struct {
	interfaces.ResourceController[*models.Device]
}

func NewDeviceController(params DeviceControllerParams) (DeviceControllerResult, error) {
	ctrl := DeviceController{}

	ctrl.ResourceController = generics.NewResourceController[*models.Device](
		params.DeviceService,
		params.LoggerService,
		params.AuthService,
		models.NewDeviceFromCreateRequest,
		models.NewDeviceFromUpdateRequest,
		generics.WithContextKey[*models.Device](deviceContextKey),
	).(*generics.ResourceController[*models.Device])

	params.Router.Mount("/devices", ctrl.ResourceController.GetRouter())

	return DeviceControllerResult{DeviceController: ctrl}, nil
}
