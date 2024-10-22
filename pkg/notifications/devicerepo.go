package notifications

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type DeviceRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService mochi.LoggerService
}

type DeviceRepoResult struct {
	fx.Out

	DeviceRepo interfaces.DeviceRepo
}

type DeviceRepo struct {
	mochi.Repository[*models.Device]

	dbService interfaces.DBService
	logger    mochi.LoggerService
}

func NewDeviceRepo(params DeviceRepoParams) DeviceRepoResult {
	embeddedRepo := mochi.NewRepository(
		params.DBService,
		params.LoggerService,
		mochi.WithTableName[*models.Device]("devices"),
	)

	repo := &DeviceRepo{
		Repository: embeddedRepo,
	}

	return DeviceRepoResult{DeviceRepo: repo}
}
