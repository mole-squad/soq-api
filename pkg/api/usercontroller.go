package api

import (
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type UserControllerParams struct {
	fx.In

	AuthService interfaces.AuthService
	UserService interfaces.UserService
	Router      *chi.Mux
}

type UserControllerResult struct {
	fx.Out

	UserController interface{}
}

type UserController struct {
	userService interfaces.UserService
}

func NewUserController(params UserControllerParams) (UserControllerResult, error) {
	ctrl := UserController{userService: params.UserService}

	userRouter := chi.NewRouter()
	userRouter.Use(params.AuthService.AuthRequired())

	params.Router.Mount("/users", userRouter)

	return UserControllerResult{
		UserController: ctrl,
	}, nil
}
