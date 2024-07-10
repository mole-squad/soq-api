package auth

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	authlib "github.com/go-pkgz/auth"
)

type AuthServiceParams struct {
	fx.In

	Logger interfaces.LoggerService
	Router *chi.Mux
}

type AuthServiceResult struct {
	fx.Out

	AuthService interfaces.AuthService
}

type AuthService struct {
	logger  interfaces.LoggerService
	authsvc *authlib.Service
}

func NewAuthService(params AuthServiceParams) (AuthServiceResult, error) {
	var result AuthServiceResult

	config := authlib.Opts{}

	authsvc := authlib.NewService(config)

	result.AuthService = &AuthService{
		logger:  params.Logger,
		authsvc: authsvc,
	}

	return result, nil
}

func (svc *AuthService) AuthRequired() func(http.Handler) http.Handler {
	middleware := svc.authsvc.Middleware()
	return middleware.Auth
}
