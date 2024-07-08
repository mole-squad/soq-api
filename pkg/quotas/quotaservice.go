package quotas

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"go.uber.org/fx"
)

type QuotaServiceParams struct {
	fx.In

	QuotaRepo interfaces.QuotaRepo
}

type QuotaServiceResult struct {
	fx.Out

	QuotaService interfaces.QuotaService
}

type QuotaService struct {
	quotaRepo interfaces.QuotaRepo
}

func NewQuotaService(params QuotaServiceParams) (QuotaServiceResult, error) {
	srv := &QuotaService{quotaRepo: params.QuotaRepo}
	return QuotaServiceResult{QuotaService: srv}, nil
}

func (srv *QuotaService) CreateUserQuota(
	ctx context.Context,
	user *models.User,
	quota *models.Quota,
) (models.Quota, error) {
	quota.UserID = user.ID

	err := srv.quotaRepo.CreateOne(ctx, quota)
	if err != nil {
		return models.Quota{}, fmt.Errorf("failed to create user quota: %w", err)
	}

	return *quota, nil
}

func (srv *QuotaService) UpdateUserQuota(
	ctx context.Context,
	quota *models.Quota,
) (models.Quota, error) {
	err := srv.quotaRepo.UpdateOne(ctx, quota)
	if err != nil {
		return models.Quota{}, fmt.Errorf("failed to update user quota: %w", err)
	}

	return *quota, nil
}

func (srv *QuotaService) DeleteUserQuota(ctx context.Context, id uint) error {
	err := srv.quotaRepo.DeleteOne(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user quota: %w", err)
	}

	return nil
}

func (srv *QuotaService) ListUserQuotas(ctx context.Context, user *models.User) ([]models.Quota, error) {
	quotas, err := srv.quotaRepo.FindManyByUser(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to list user quotas: %w", err)
	}

	return quotas, nil
}
