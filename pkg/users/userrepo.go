package users

import (
	"context"
	"fmt"

	"github.com/burkel24/go-mochi"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type UserRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService mochi.LoggerService
}

type UserRepoResult struct {
	fx.Out

	UserRepo interfaces.UserRepo
}

type UserRepo struct {
	dbService interfaces.DBService
	logger    mochi.LoggerService
}

func NewUserRepo(params UserRepoParams) (UserRepoResult, error) {
	repo := &UserRepo{
		dbService: params.DBService,
		logger:    params.LoggerService,
	}

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

func (repo *UserRepo) CreateOne(ctx context.Context, user *models.User) error {
	repo.logger.Debug("Creating one user", "user", user.Username)

	err := repo.dbService.CreateOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
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
	err := repo.dbService.DEPUpdateOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
