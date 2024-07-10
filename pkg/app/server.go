package app

import (
	"context"
	"net"
	"net/http"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

const port = ":3000"

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
