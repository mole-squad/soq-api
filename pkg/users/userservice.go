package users

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
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

func (srv *UserService) UpdateUserPassword(ctx context.Context, userID uint, password string) error {
	hashedPassword, err := HashUserPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		Model:        gorm.Model{ID: userID},
		PasswordHash: hashedPassword,
	}

	err = srv.userRepo.UpdateOne(ctx, &user)
	if err != nil {
		return fmt.Errorf("failed to update user password: %w", err)
	}

	return nil
}
