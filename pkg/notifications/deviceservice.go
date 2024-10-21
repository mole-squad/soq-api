package notifications

import (
	"github.com/burkel24/go-mochi"
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
	mochi.Service[*models.Device]
}

func NewDeviceService(params DeviceServiceParams) DeviceServiceResult {
	embeddedSvc := mochi.NewService(
		params.DeviceRepo,
	)

	srv := &DeviceService{
		Service: embeddedSvc,
	}

	return DeviceServiceResult{DeviceService: srv}
}
