package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mole-squad/soq/pkg/interfaces"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, router *chi.Mux, logger interfaces.LoggerService) *http.Server {
	portStr := os.Getenv("PORT")
	port := fmt.Sprintf(":%s", portStr)

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
