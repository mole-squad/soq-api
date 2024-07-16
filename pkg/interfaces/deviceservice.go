package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type DeviceService interface {
	CreateUserDevice(ctx context.Context, user *models.User, device *models.Device) (models.Device, error)
	GetUserDevice(ctx context.Context, userID uint, deviceID uint) (models.Device, error)
	UpdateUserDevice(ctx context.Context, device *models.Device) (models.Device, error)
	DeleteUserDevice(ctx context.Context, id uint) error
	ListUserDevices(ctx context.Context, userID uint) ([]models.Device, error)
}
