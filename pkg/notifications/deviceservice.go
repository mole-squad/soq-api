package notifications

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type DeviceServiceParams struct {
	fx.In

	DeviceRepo    interfaces.DeviceRepo
	LoggerService interfaces.LoggerService
}

type DeviceServiceResult struct {
	fx.Out

	DeviceService interfaces.DeviceService
}

type DeviceService struct {
	deviceRepo interfaces.DeviceRepo
	logger     interfaces.LoggerService
}

func NewDeviceService(p DeviceServiceParams) DeviceServiceResult {
	srv := &DeviceService{
		deviceRepo: p.DeviceRepo,
		logger:     p.LoggerService,
	}

	return DeviceServiceResult{DeviceService: srv}
}

func (srv *DeviceService) CreateUserDevice(ctx context.Context, user *models.User, device *models.Device) (models.Device, error) {
	device.UserID = user.ID

	err := srv.deviceRepo.CreateOne(ctx, device)
	if err != nil {
		return models.Device{}, err
	}

	return *device, nil
}

func (srv *DeviceService) GetUserDevice(ctx context.Context, userID uint, deviceID uint) (models.Device, error) {
	device, err := srv.deviceRepo.FindOneByUser(ctx, userID, "devices.id = ?", deviceID)
	if err != nil {
		return models.Device{}, fmt.Errorf("failed to get user device: %w", err)
	}

	return *device, nil
}

func (srv *DeviceService) UpdateUserDevice(ctx context.Context, device *models.Device) (models.Device, error) {
	err := srv.deviceRepo.UpdateOne(ctx, device)
	if err != nil {
		return models.Device{}, fmt.Errorf("failed to update user device: %w", err)
	}

	return *device, nil
}

func (srv *DeviceService) DeleteUserDevice(ctx context.Context, id uint) error {
	err := srv.deviceRepo.DeleteOne(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user device: %w", err)
	}

	return nil
}

func (srv *DeviceService) ListUserDevices(ctx context.Context, userID uint) ([]models.Device, error) {
	devices, err := srv.deviceRepo.FindManyByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list user devices: %w", err)
	}

	return devices, nil
}
