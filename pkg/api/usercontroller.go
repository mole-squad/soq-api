package api

import (
	"net/http"

	"github.com/burkel24/go-mochi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/api"
	"github.com/mole-squad/soq-api/pkg/common"
	"github.com/mole-squad/soq-api/pkg/models"
	"github.com/mole-squad/soq-api/pkg/users"
	"go.uber.org/fx"
)

type UserControllerParams struct {
	fx.In

	AuthService mochi.AuthService
	UserService mochi.UserService
	Router      *chi.Mux
}

type UserControllerResult struct {
	fx.Out

	UserController interface{}
}

type UserController struct {
	auth        mochi.AuthService
	userService mochi.UserService
}

func NewUserController(params UserControllerParams) (UserControllerResult, error) {
	ctrl := UserController{
		auth:        params.AuthService,
		userService: params.UserService,
	}

	userRouter := chi.NewRouter()
	userRouter.Use(params.AuthService.AuthRequired())

	userRouter.Patch("/password", ctrl.SetPassword)

	userRouter.With(params.AuthService.AdminRequired()).Post("/", ctrl.CreateUser)

	params.Router.Mount("/users", userRouter)

	return UserControllerResult{
		UserController: ctrl,
	}, nil
}

func (ctrl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto := &api.CreateUserRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	hashedPassword, err := users.HashUserPassword(dto.Password)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
	}

	newUser := &models.User{
		Username:     dto.Username,
		Name:         dto.Name,
		PasswordHash: hashedPassword,
		Timezone:     dto.Timezone,
	}

	user, err := ctrl.userService.CreateUser(ctx, newUser)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, user.(*models.User).ToDTO())

}

func (ctrl *UserController) SetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := ctrl.auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	dto := &users.SetPasswordRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	err = ctrl.userService.UpdateUserPassword(ctx, user.GetID(), dto.Password)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
