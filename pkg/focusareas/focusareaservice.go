package focusareas

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"go.uber.org/fx"
)

type FocusAreaServiceParams struct {
	fx.In

	FocusAreaRepo interfaces.FocusAreaRepo
}

type FocusAreaServiceResult struct {
	fx.Out

	FocusAreaService interfaces.FocusAreaService
}

type FocusAreaService struct {
	focusAreaRepo interfaces.FocusAreaRepo
}

func NewFocusAreaService(params FocusAreaServiceParams) (FocusAreaServiceResult, error) {
	srv := &FocusAreaService{focusAreaRepo: params.FocusAreaRepo}
	return FocusAreaServiceResult{FocusAreaService: srv}, nil
}

func (srv *FocusAreaService) CreateFocusArea(ctx context.Context, user *models.User, focusArea *models.FocusArea) (models.FocusArea, error) {
	focusArea.UserID = user.ID

	err := srv.focusAreaRepo.CreateOne(ctx, focusArea)
	if err != nil {
		return models.FocusArea{}, fmt.Errorf("failed to create focus area: %w", err)
	}

	return *focusArea, nil
}

func (srv *FocusAreaService) UpdateFocusArea(ctx context.Context, focusArea *models.FocusArea) (models.FocusArea, error) {
	err := srv.focusAreaRepo.UpdateOne(ctx, focusArea)
	if err != nil {
		return models.FocusArea{}, fmt.Errorf("failed to update focus area: %w", err)
	}

	return *focusArea, nil
}

func (srv *FocusAreaService) DeleteFocusArea(ctx context.Context, id uint) error {
	err := srv.focusAreaRepo.DeleteOne(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete focus area: %w", err)
	}

	return nil
}

func (srv *FocusAreaService) ListFocusAreas(ctx context.Context, user *models.User) ([]models.FocusArea, error) {
	focusAreas, err := srv.focusAreaRepo.FindManyByUser(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to list focus areas: %w", err)
	}

	return focusAreas, nil
}
