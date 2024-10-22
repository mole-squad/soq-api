package auth

import (
	"github.com/burkel24/go-mochi"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"Auth",
	fx.Provide(mochi.NewAuthService),
)
