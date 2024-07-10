package app

import (
	"log/slog"
	"net/http"

	"github.com/burkel24/task-app/pkg/agendas"
	"github.com/burkel24/task-app/pkg/api"
	"github.com/burkel24/task-app/pkg/auth"
	"github.com/burkel24/task-app/pkg/db"
	"github.com/burkel24/task-app/pkg/focusareas"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/logger"
	"github.com/burkel24/task-app/pkg/notifications"
	"github.com/burkel24/task-app/pkg/quotas"
	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/burkel24/task-app/pkg/users"
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
