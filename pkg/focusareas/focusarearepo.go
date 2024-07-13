package focusareas

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type FocusAreaRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type FocusAreaRepoResult struct {
	fx.Out

	FocusAreaRepo interfaces.FocusAreaRepo
}

type FocusAreaRepo struct {
	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewFocusAreaRepo(params FocusAreaRepoParams) (FocusAreaRepoResult, error) {
	repo := &FocusAreaRepo{
		dbService: params.DBService,
		logger:    params.LoggerService,
	}

	return FocusAreaRepoResult{FocusAreaRepo: repo}, nil
}

func (repo *FocusAreaRepo) CreateOne(ctx context.Context, focusArea *models.FocusArea) error {
	repo.logger.Info("Creating one focus area", "focusArea", focusArea)

	err := repo.dbService.CreateOne(ctx, focusArea)
	if err != nil {
		return fmt.Errorf("failed to create one focus area: %w", err)
	}

	return nil
}

func (repo *FocusAreaRepo) UpdateOne(ctx context.Context, focusArea *models.FocusArea) error {
	repo.logger.Info("Updating one focus area", "focusArea", focusArea)

	err := repo.dbService.UpdateOne(ctx, focusArea)
	if err != nil {
		return fmt.Errorf("failed to update one focus area: %w", err)
	}

	return nil
}

func (repo *FocusAreaRepo) DeleteOne(ctx context.Context, id uint) error {
	repo.logger.Info("Deleting one focus area", "id", id)

	focusArea := &models.FocusArea{Model: gorm.Model{ID: id}}

	err := repo.dbService.DeleteOne(ctx, focusArea)
	if err != nil {
		return fmt.Errorf("failed to delete one focus area: %w", err)
	}

	return nil
}

func (repo *FocusAreaRepo) FindManyByUser(ctx context.Context, userID uint) ([]models.FocusArea, error) {
	repo.logger.Info("Finding many focus areas by user", "userID", userID)

	focusAreas := []models.FocusArea{}

	err := repo.dbService.FindMany(ctx, &focusAreas, []string{}, []string{"TimeWindows"}, "user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find many focus areas by user: %w", err)
	}

	return focusAreas, nil
}
