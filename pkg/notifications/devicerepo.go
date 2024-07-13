package notifications

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type DeviceRepoParams struct {
	fx.In

	DbService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type DeviceRepoResult struct {
	fx.Out

	DeviceRepo interfaces.DeviceRepo
}

type DeviceRepo struct {
	dbService     interfaces.DBService
	loggerService interfaces.LoggerService
}

func NewDeviceRepo(p DeviceRepoParams) DeviceRepoResult {
	repo := &DeviceRepo{
		dbService:     p.DbService,
		loggerService: p.LoggerService,
	}

	return DeviceRepoResult{DeviceRepo: repo}
}

func (r *DeviceRepo) FindManyByUser(ctx context.Context, userID uint) ([]models.Device, error) {
	var devices []models.Device

	err := r.dbService.FindMany(ctx, &devices, []string{}, []string{}, "user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find devices by user: %w", err)
	}

	return devices, nil
}
