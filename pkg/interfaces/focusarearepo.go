package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type FocusAreaRepo interface {
	CreateOne(ctx context.Context, focusArea *models.FocusArea) error
	UpdateOne(ctx context.Context, focusArea *models.FocusArea) error
	DeleteOne(ctx context.Context, id uint) error
	FindManyByUser(ctx context.Context, userID uint) ([]models.FocusArea, error)
}
