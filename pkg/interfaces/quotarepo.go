package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type QuotaRepo interface {
	CreateOne(ctx context.Context, quota *models.Quota) error
	UpdateOne(ctx context.Context, quota *models.Quota) error
	DeleteOne(ctx context.Context, id uint) error
	FindManyByUser(ctx context.Context, userID uint) ([]models.Quota, error)
}
