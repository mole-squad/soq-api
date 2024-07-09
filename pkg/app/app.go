package app

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	"github.com/burkel24/task-app/pkg/agendas"
	"github.com/burkel24/task-app/pkg/db"
	"github.com/burkel24/task-app/pkg/focusareas"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/logger"
	"github.com/burkel24/task-app/pkg/quotas"
	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/burkel24/task-app/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

const port = ":3000"

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("okay xD"))
	})

	return router
}

func NewServer(lc fx.Lifecycle, router *chi.Mux, logger interfaces.LoggerService) *http.Server {
	srv := &http.Server{Addr: port, Handler: router}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			logger.Info("Starting HTTP server", "port", srv.Addr)
			go srv.Serve(ln)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down HTTP server")

			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func NewFxLogger(logger interfaces.LoggerService) fxevent.Logger {
	fxLogger := fxevent.SlogLogger{Logger: logger.Logger()}

	fxLogger.UseLogLevel(slog.LevelDebug)
	fxLogger.UseErrorLevel(slog.LevelError)

	return &fxLogger
}

func BuildServerOpts() []fx.Option {
	return []fx.Option{
		fx.Provide(NewServer),
		fx.Invoke(func(*http.Server) {}),
	}
}

func BuildAppOpts() []fx.Option {
	return []fx.Option{
		fx.WithLogger(NewFxLogger),
		fx.Provide(NewRouter),
		fx.Provide(logger.NewLoggerService),
		db.Module,
		users.Module,
		focusareas.Module,
		tasks.Module,
		quotas.Module,
		agendas.Module,
	}
}
