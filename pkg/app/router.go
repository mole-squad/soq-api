package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("okay xD"))
	})

	return router
}
