package interfaces

import (
	"context"
	"time"

	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/models"
)

type AgendaRepo interface {
	mochi.Repository[*models.Agenda]

	FindManyByStatus(ctx context.Context, status models.AgendaStatus) ([]models.Agenda, error)
	FindOneByTimeRangeFocusArea(
		ctx context.Context,
		userID uint,
		focusAreaID uint,
		startTime, endTime time.Time,
	) (*models.Agenda, error)
}
