package app

import (
	"log/slog"
	"net/http"

	"github.com/mole-squad/soq-api/pkg/agendas"
	"github.com/mole-squad/soq-api/pkg/api"
	"github.com/mole-squad/soq-api/pkg/auth"
	"github.com/mole-squad/soq-api/pkg/db"
	"github.com/mole-squad/soq-api/pkg/focusareas"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/logger"
	"github.com/mole-squad/soq-api/pkg/notifications"
	"github.com/mole-squad/soq-api/pkg/quotas"
	"github.com/mole-squad/soq-api/pkg/tasks"
	"github.com/mole-squad/soq-api/pkg/users"
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
