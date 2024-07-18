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
	*generics.ResourceService[*models.Device]
}

func NewDeviceService(params DeviceServiceParams) DeviceServiceResult {
	embeddedSvc := generics.NewResourceService[*models.Device](
		params.DeviceRepo,
	).(*generics.ResourceService[*models.Device])

	srv := &DeviceService{
		ResourceService: embeddedSvc,
	}

	return DeviceServiceResult{DeviceService: srv}
}
