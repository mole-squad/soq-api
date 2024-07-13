package interfaces

import (
	"context"

	"github.com/mole-squad/soq-api/pkg/models"
)

type FocusAreaRepo interface {
	CreateOne(ctx context.Context, focusArea *models.FocusArea) error
	UpdateOne(ctx context.Context, focusArea *models.FocusArea) error
	DeleteOne(ctx context.Context, id uint) error
	FindManyByUser(ctx context.Context, userID uint) ([]models.FocusArea, error)
}
