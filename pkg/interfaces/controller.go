package interfaces

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Controller[M Resource] interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)

	ItemFromContext(ctx context.Context) (M, error)
	ItemContextMiddleware(next http.Handler) http.Handler
	UserAccessMiddleware(next http.Handler) http.Handler

	GetRouter() *chi.Mux
}
