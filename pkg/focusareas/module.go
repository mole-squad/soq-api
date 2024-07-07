package focusareas

import "go.uber.org/fx"

var Module = fx.Module(
	"FocusAreas",
	fx.Provide(fx.Private, NewFocusAreaRepo),
	fx.Invoke(NewFocusAreaController),
	fx.Provide(NewFocusAreaService),
)
