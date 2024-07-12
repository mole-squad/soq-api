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

func (srv *UserService) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	user, err := srv.userRepo.FindOneByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (srv *UserService) GetUserByCredentials(ctx context.Context, username, passwordHash string) (*models.User, error) {
	user, err := srv.userRepo.FindOneByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	if DoPasswordsMatch(user.PasswordHash, passwordHash) {
		return user, nil
	}

	return nil, fmt.Errorf("invalid password")
}
