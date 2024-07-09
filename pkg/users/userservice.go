package users

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
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
	return UserServiceResult{UserService: &srv}, nil
}

func (srv *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	users, err := srv.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}
