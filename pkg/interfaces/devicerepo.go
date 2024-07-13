package interfaces

import (
	"context"

	"github.com/mole-squad/soq/pkg/models"
)

type DeviceRepo interface {
	FindManyByUser(ctx context.Context, userID uint) ([]models.Device, error)
}
