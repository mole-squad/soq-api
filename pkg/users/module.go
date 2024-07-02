package users

import "go.uber.org/fx"

var Module = fx.Module(
	"User",
	fx.Provide(fx.Private, NewUserRepo),
	fx.Invoke(NewUserController),
	fx.Provide(NewUserService),
)
