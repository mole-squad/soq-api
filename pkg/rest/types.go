package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
)

type ResourceContextKey int

type Resource interface {
	GetID() uint
	ToDTO() render.Renderer
}

type CRUDService[M Resource] interface {
	List(ctx context.Context, userID uint) ([]M, error)
	CreateOne(ctx context.Context, userID uint, item M) (M, error)
	GetOne(ctx context.Context, userID, itemID uint) (M, error)
	UpdateOne(ctx context.Context, userID uint, itemID uint, item M) (M, error)
	DeleteOne(ctx context.Context, userID uint, itemID uint) error
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type ResourceRequestConstructor[M Resource] func(r *http.Request) (M, error)
