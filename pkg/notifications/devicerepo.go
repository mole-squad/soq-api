package notifications

import (
	"github.com/mole-squad/soq-api/pkg/generics"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type DeviceRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type DeviceRepoResult struct {
	fx.Out

	DeviceRepo interfaces.DeviceRepo
}

type DeviceRepo struct {
	*generics.ResourceRepository[*models.Device]

	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewDeviceRepo(params DeviceRepoParams) DeviceRepoResult {
	embeddedRepo := generics.NewResourceRepository[*models.Device](
		params.DBService,
		params.LoggerService,
		generics.WithTableName[*models.Device]("devices"),
	).(*generics.ResourceRepository[*models.Device])

	repo := &DeviceRepo{
		ResourceRepository: embeddedRepo,
	}

	return DeviceRepoResult{DeviceRepo: repo}
}
