package notifications

import "go.uber.org/fx"

var Module = fx.Module(
	"Notifications",
	fx.Provide(fx.Private, NewDeviceRepo),
	fx.Provide(NewDeviceService),
	fx.Provide(NewNotificationService),
	fx.Provide(NewDeviceRepo),
)
