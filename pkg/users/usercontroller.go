package users

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type UserControllerParams struct {
	fx.In

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
	userRouter.Get("/", ctrl.ListUsers)

	params.Router.Mount("/users", userRouter)

	return UserControllerResult{UserController: ctrl}, nil
}

func (ctrl *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("these are my users xD"))
}
