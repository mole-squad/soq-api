package tasks

import "go.uber.org/fx"

var Module = fx.Module(
	"Tasks",
	fx.Provide(fx.Private, NewTaskRepo),
	fx.Provide(NewTaskService),
)
