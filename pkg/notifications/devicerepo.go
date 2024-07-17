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
	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewDeviceRepo(p DeviceRepoParams) DeviceRepoResult {
	repo := &DeviceRepo{
		dbService: p.DbService,
		logger:    p.LoggerService,
	}

	return DeviceRepoResult{DeviceRepo: repo}
}

func (r *DeviceRepo) CreateOne(ctx context.Context, device *models.Device) error {
	err := r.dbService.CreateOne(ctx, device)
	if err != nil {
		return fmt.Errorf("failed to create device: %w", err)
	}

	r.logger.Debug("Created device", "device", device.ID, "user", device.UserID)

	return nil
}

func (r *DeviceRepo) FindOneByUser(ctx context.Context, userID uint, query string, args ...interface{}) (*models.Device, error) {
	var device models.Device

	fullQuery := "devices.user_id = ?"
	if query != "" {
		fullQuery = fmt.Sprintf("%s AND %s", fullQuery, query)
	}

	fullArgs := append([]interface{}{userID}, args...)

	err := r.dbService.FindOne(ctx, &device, []string{}, []string{}, fullQuery, fullArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to find device: %w", err)
	}

	return &device, nil

}

func (r *DeviceRepo) UpdateOne(ctx context.Context, device *models.Device) error {
	err := r.dbService.UpdateOne(ctx, device)
	if err != nil {
		return fmt.Errorf("failed to update device: %w", err)
	}

	r.logger.Debug("Updated device", "device", device.ID, "user", device.UserID)

	return nil
}

func (r *DeviceRepo) DeleteOne(ctx context.Context, id uint) error {
	device := &models.Device{}

	err := r.dbService.DeleteOne(ctx, id, device)
	if err != nil {
		return fmt.Errorf("failed to delete device: %w", err)
	}

	r.logger.Debug("Deleted device", "device", id)

	return nil
}

func (r *DeviceRepo) FindManyByUser(ctx context.Context, userID uint) ([]models.Device, error) {
	var devices []models.Device

	err := r.dbService.FindMany(ctx, &devices, []string{}, []string{}, "user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find devices by user: %w", err)
	}

	return devices, nil
}
