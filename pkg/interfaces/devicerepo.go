package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type DeviceRepo interface {
	FindManyByUser(ctx context.Context, userID uint) ([]models.Device, error)
}
