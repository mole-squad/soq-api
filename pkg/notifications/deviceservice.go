package notifications

import (
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type DeviceServiceParams struct {
	fx.In

	DeviceRepo interfaces.DeviceRepo
}

type DeviceServiceResult struct {
	fx.Out

	DeviceService interfaces.DeviceService
}

type DeviceService struct {
	*generics.Service[*models.Device]
}

func NewDeviceService(params DeviceServiceParams) DeviceServiceResult {
	embeddedSvc := generics.NewService[*models.Device](
		params.DeviceRepo,
	).(*generics.Service[*models.Device])

	srv := &DeviceService{
		Service: embeddedSvc,
	}

	return DeviceServiceResult{DeviceService: srv}
}
