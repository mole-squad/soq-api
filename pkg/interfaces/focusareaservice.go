package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type FocusAreaService interface {
	CreateFocusArea(ctx context.Context, user *models.User, focusArea *models.FocusArea) (models.FocusArea, error)
	UpdateFocusArea(ctx context.Context, focusArea *models.FocusArea) (models.FocusArea, error)
	DeleteFocusArea(ctx context.Context, id uint) error
	ListUserFocusAreas(ctx context.Context, user *models.User) ([]models.FocusArea, error)
}
