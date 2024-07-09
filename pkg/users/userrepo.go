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
