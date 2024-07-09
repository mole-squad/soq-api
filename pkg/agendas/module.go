package agendas

import "go.uber.org/fx"

var Module = fx.Module(
	"Agendas",
	fx.Provide(fx.Private, NewAgendaRepo),
	fx.Provide(NewAgendaService),
)
