package users

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"go.uber.org/fx"
)

type UserRepoParams struct {
	fx.In

	DBService interfaces.DBService
}

type UserRepoResult struct {
	fx.Out

	UserRepo interfaces.UserRepo
}

type UserRepo struct {
	dbService interfaces.DBService
}

func NewUserRepo(params UserRepoParams) (UserRepoResult, error) {
	repo := &UserRepo{dbService: params.DBService}

	return UserRepoResult{UserRepo: repo}, nil
}

func (repo *UserRepo) ListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	err := repo.dbService.FindMany(ctx, &users, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}

func (repo *UserRepo) FindOneByID(ctx context.Context, userID uint) (*models.User, error) {
	var user models.User

	err := repo.dbService.FindOne(ctx, &user, []string{}, []string{}, "id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &user, nil
}

func (repo *UserRepo) FindOneByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := repo.dbService.FindOne(ctx, &user, []string{}, []string{}, "username = ?", username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return &user, nil
}

func (repo *UserRepo) UpdateOne(ctx context.Context, user *models.User) error {
	err := repo.dbService.UpdateOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
