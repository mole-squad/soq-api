package users

import (
	"github.com/burkel24/task-app/pkg/interfaces"
	"go.uber.org/fx"
)

type UserServiceParams struct {
	fx.In

	UserRepo interfaces.UserRepo
}

type UserServiceResult struct {
	fx.Out

	UserService interfaces.UserService
}

type UserService struct {
	userRepo interfaces.UserRepo
}

func NewUserService(params UserServiceParams) (UserServiceResult, error) {
	srv := UserService{userRepo: params.UserRepo}
	return UserServiceResult{UserService: srv}, nil
}
