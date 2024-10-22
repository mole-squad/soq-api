package interfaces

import (
	"context"

	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/models"
)

type AgendaService interface {
	mochi.Service[*models.Agenda]
	GenerateAgendasForUpcomingTimeWindows(ctx context.Context) error
	PopulatePendingAgendas(ctx context.Context) error
	SendAgendaNotifications(ctx context.Context) error
}
