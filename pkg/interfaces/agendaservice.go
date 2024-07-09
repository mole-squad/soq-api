package interfaces

import "context"

type AgendaService interface {
	GenerateAgendasForUpcomingTimeWindows(ctx context.Context) error
	PopulatePendingAgendas(ctx context.Context) error
}
