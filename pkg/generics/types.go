package generics

import (
	"net/http"

	"github.com/mole-squad/soq-api/pkg/interfaces"
)

type ResourceContextKey int

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type ResourceRequestConstructor[M interfaces.Resource] func(r *http.Request) (M, error)

type Query struct {
	Filter string
	Args   []interface{}
}
