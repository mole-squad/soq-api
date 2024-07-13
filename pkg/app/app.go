package app

import (
	"log/slog"
	"net/http"

	"github.com/mole-squad/soq/pkg/agendas"
	"github.com/mole-squad/soq/pkg/api"
	"github.com/mole-squad/soq/pkg/auth"
	"github.com/mole-squad/soq/pkg/db"
	"github.com/mole-squad/soq/pkg/focusareas"
	"github.com/mole-squad/soq/pkg/interfaces"
	"github.com/mole-squad/soq/pkg/logger"
	"github.com/mole-squad/soq/pkg/notifications"
	"github.com/mole-squad/soq/pkg/quotas"
	"github.com/mole-squad/soq/pkg/tasks"
	"github.com/mole-squad/soq/pkg/users"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewFxLogger(logger interfaces.LoggerService) fxevent.Logger {
	fxLogger := fxevent.SlogLogger{Logger: logger.Logger()}

	fxLogger.UseLogLevel(slog.LevelDebug)
	fxLogger.UseErrorLevel(slog.LevelError)

	return &fxLogger
}

func BuildServerOpts() []fx.Option {
	return []fx.Option{
		fx.Provide(NewRouter),
		fx.Provide(NewServer),
		fx.Invoke(func(*http.Server) {}),
		auth.Module,
		api.Module,
	}
}

func BuildAppOpts() []fx.Option {
	return []fx.Option{
		fx.WithLogger(NewFxLogger),
		fx.Provide(logger.NewLoggerService),

		db.Module,
		users.Module,
		focusareas.Module,
		notifications.Module,
		tasks.Module,
		quotas.Module,
		agendas.Module,
	}
}
