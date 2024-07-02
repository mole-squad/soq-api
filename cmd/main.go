package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/burkel24/task-app/pkg/db"
	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/burkel24/task-app/pkg/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

const port = ":3000"

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("okay xD"))
	})

	return router
}

func NewServer(lc fx.Lifecycle, router *chi.Mux) *http.Server {

	srv := &http.Server{Addr: port, Handler: router}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down")
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func main() {
	_, err := db.InitDb()

	if err != nil {
		slog.Error("DB connection failed %w", slog.Attr{Key: "error", Value: slog.AnyValue(err)})
		os.Exit(1)
	}

	fx.New(
		db.Module,
		users.Module,
		tasks.Module,
		fx.Provide(NewRouter),
		fx.Provide(NewServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
