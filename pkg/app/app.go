package app

import (
	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/agendas"
	"github.com/mole-squad/soq-api/pkg/api"
	"github.com/mole-squad/soq-api/pkg/db"
	"github.com/mole-squad/soq-api/pkg/focusareas"
	"github.com/mole-squad/soq-api/pkg/notifications"
	"github.com/mole-squad/soq-api/pkg/quotas"
	"github.com/mole-squad/soq-api/pkg/tasks"
	"github.com/mole-squad/soq-api/pkg/timewindows"
	"github.com/mole-squad/soq-api/pkg/users"
	"go.uber.org/fx"
)

func BuildServerOpts() []fx.Option {
	return append(mochi.BuildServerOpts(), []fx.Option{
		api.Module,
	}...)
}

func BuildAppOpts() []fx.Option {
	return append(mochi.BuildAppOpts(), []fx.Option{
		db.Module,
		users.Module,
		focusareas.Module,
		notifications.Module,
		tasks.Module,
		timewindows.Module,
		quotas.Module,
		agendas.Module,
	}...)
}
