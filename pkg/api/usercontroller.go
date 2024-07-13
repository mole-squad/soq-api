package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/pkg/auth"
	"github.com/mole-squad/soq-api/pkg/common"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/users"
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

	userRouter.Patch("/password", ctrl.SetPassword)

	params.Router.Mount("/users", userRouter)

	return UserControllerResult{
		UserController: ctrl,
	}, nil
}

func (ctrl *UserController) SetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	dto := &users.SetPasswordRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	err = ctrl.userService.UpdateUserPassword(ctx, user.ID, dto.Password)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
