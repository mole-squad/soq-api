package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type DeviceRepo interface {
	CreateOne(ctx context.Context, device *models.Device) error
	FindOneByUser(ctx context.Context, userID uint, query string, args ...interface{}) (*models.Device, error)
	UpdateOne(ctx context.Context, device *models.Device) error
	DeleteOne(ctx context.Context, id uint) error
	FindManyByUser(ctx context.Context, userID uint) ([]models.Device, error)
}
