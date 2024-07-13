package interfaces

import (
	"context"
	"time"

	"github.com/mole-squad/soq/pkg/models"
)

type AgendaRepo interface {
	CreateOne(ctx context.Context, agenda *models.Agenda) error
	UpdateOne(ctx context.Context, agenda *models.Agenda) error
	FindManyByUser(ctx context.Context, userID uint) ([]models.Agenda, error)
	FindManyByStatus(ctx context.Context, status models.AgendaStatus) ([]models.Agenda, error)
	FindOneByTimeRangeFocusArea(
		ctx context.Context,
		userID uint,
		focusAreaID uint,
		startTime, endTime time.Time,
	) (*models.Agenda, error)
}
