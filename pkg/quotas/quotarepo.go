package quotas

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type QuotaRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type QuotaRepoResult struct {
	fx.Out

	QuotaRepo interfaces.QuotaRepo
}

type QuotaRepo struct {
	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewQuotaRepo(params QuotaRepoParams) (QuotaRepoResult, error) {
	repo := &QuotaRepo{
		dbService: params.DBService,
		logger:    params.LoggerService,
	}

	return QuotaRepoResult{QuotaRepo: repo}, nil
}

func (repo *QuotaRepo) CreateOne(ctx context.Context, quota *models.Quota) error {
	repo.logger.Info("Creating one quota", "quota", quota)

	err := repo.dbService.CreateOne(ctx, quota)
	if err != nil {
		return fmt.Errorf("failed to create one quota: %w", err)
	}

	return nil
}

func (repo *QuotaRepo) UpdateOne(ctx context.Context, quota *models.Quota) error {
	repo.logger.Info("Updating one quota", "quota", quota)

	err := repo.dbService.UpdateOne(ctx, quota)
	if err != nil {
		return fmt.Errorf("failed to update one quota: %w", err)
	}

	return nil
}

func (repo *QuotaRepo) DeleteOne(ctx context.Context, id uint) error {
	repo.logger.Info("Deleting one quota", "id", id)

	quota := &models.Quota{Model: gorm.Model{ID: id}}
	err := repo.dbService.DeleteOne(ctx, quota)
	if err != nil {
		return fmt.Errorf("failed to delete one quota: %w", err)
	}

	return nil
}

func (repo *QuotaRepo) FindManyByUser(ctx context.Context, userID uint) ([]models.Quota, error) {
	var quotas []models.Quota

	err := repo.dbService.FindMany(ctx, &quotas, []string{"FocusArea"}, []string{}, "quota.user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find many quotas by user: %w", err)
	}

	return quotas, nil
}
