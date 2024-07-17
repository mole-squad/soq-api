package interfaces

import (
	"context"
	"net/http"

	"github.com/mole-squad/soq-api/pkg/models"
)

type AuthService interface {
	AuthRequired() func(http.Handler) http.Handler
	AdminRequired() func(http.Handler) http.Handler
	GetUserFromCtx(ctx context.Context) (*models.User, error)
	LoginUser(ctx context.Context, username, password string) (string, error)
}
