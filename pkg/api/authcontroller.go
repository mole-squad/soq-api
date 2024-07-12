package api

import (
	"fmt"
	"net/http"

	"github.com/burkel24/task-app/pkg/auth"
	"github.com/burkel24/task-app/pkg/common"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

type AuthControllerParams struct {
	fx.In

	AuthService interfaces.AuthService
	Router      *chi.Mux
}

type AuthControllerResult struct {
	fx.Out

	AuthController AuthController
}

type AuthController struct {
	authService interfaces.AuthService
}

func NewAuthController(params AuthControllerParams) (AuthControllerResult, error) {
	ctrl := AuthController{authService: params.AuthService}

	authRouter := chi.NewRouter()

	authRouter.Post("/token", ctrl.GetToken)

	params.Router.Mount("/auth", authRouter)

	return AuthControllerResult{AuthController: ctrl}, nil
}

func (ctrl *AuthController) GetToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto := &auth.LoginRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	token, err := ctrl.authService.LoginUser(ctx, dto.Username, dto.Password)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(fmt.Errorf("invalid credentials")))
		return
	}

	resp := auth.NewTokenResponseDTO(token)
	render.Render(w, r, resp)
}