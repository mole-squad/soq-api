package interfaces

import (
	"context"
	"net/http"
)

type AuthService interface {
	AuthRequired() func(http.Handler) http.Handler
	LoginUser(ctx context.Context, username, password string) (string, error)
}
