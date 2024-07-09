package interfaces

import (
	"context"
	"time"

	"github.com/burkel24/task-app/pkg/models"
)

type AgendaRepo interface {
	CreateOne(ctx context.Context, agenda *models.Agenda) error
	FindManyByUser(ctx context.Context, userID uint) ([]models.Agenda, error)
	FindManyByPending(ctx context.Context) ([]models.Agenda, error)
	FindOneByTimeRangeFocusArea(
		ctx context.Context,
		userID uint,
		focusAreaID uint,
		startTime, endTime time.Time,
	) (*models.Agenda, error)
}
