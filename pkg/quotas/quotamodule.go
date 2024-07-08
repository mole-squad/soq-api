package quotas

import "go.uber.org/fx"

var Module = fx.Module(
	"Quotas",
	fx.Provide(fx.Private, NewQuotaRepo),
	fx.Invoke(NewQuotaController),
	fx.Provide(NewQuotaService),
)
