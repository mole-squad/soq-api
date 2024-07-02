package users

import (
	"github.com/burkel24/task-app/pkg/interfaces"
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

func NewUserRepo(params UserRepoParams) (interfaces.UserRepo, error) {
	repo := UserRepo{dbService: params.DBService}

	return UserRepoResult{UserRepo: repo}, nil
}
