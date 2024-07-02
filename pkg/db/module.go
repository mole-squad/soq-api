package db

import "go.uber.org/fx"

var Module = fx.Module(
	"DB",
	fx.Provide(NewDBService),
)
