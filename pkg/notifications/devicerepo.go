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
	*generics.Repository[*models.Device]

	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewDeviceRepo(params DeviceRepoParams) DeviceRepoResult {
	embeddedRepo := generics.NewRepository[*models.Device](
		params.DBService,
		params.LoggerService,
		generics.WithTableName[*models.Device]("devices"),
	).(*generics.Repository[*models.Device])

	repo := &DeviceRepo{
		Repository: embeddedRepo,
	}

	return DeviceRepoResult{DeviceRepo: repo}
}
