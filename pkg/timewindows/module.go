package timewindows

import "go.uber.org/fx"

var Module = fx.Module(
	"TimeWindows",
	fx.Provide(fx.Private, NewTimeWindowRepo),
	fx.Provide(NewTimeWindowService),
)
