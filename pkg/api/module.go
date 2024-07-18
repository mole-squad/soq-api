package api

import "go.uber.org/fx"

var Module = fx.Module(
	"API",
	fx.Invoke(NewAuthController),
	fx.Invoke(NewFocusAreaController),
	fx.Invoke(NewDeviceController),
	fx.Invoke(NewQuotaController),
	fx.Invoke(NewTaskController),
	fx.Invoke(NewTimeWindowController),
	fx.Invoke(NewUserController),
)
