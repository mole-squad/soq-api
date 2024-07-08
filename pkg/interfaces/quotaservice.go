package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type QuotaService interface {
	CreateUserQuota(ctx context.Context, user *models.User, quota *models.Quota) (models.Quota, error)
	UpdateUserQuota(ctx context.Context, quota *models.Quota) (models.Quota, error)
	DeleteUserQuota(ctx context.Context, id uint) error
	ListUserQuotas(ctx context.Context, user *models.User) ([]models.Quota, error)
}
