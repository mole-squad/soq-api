package interfaces

import (
	"context"

	"github.com/mole-squad/soq/pkg/models"
)

type UserRepo interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	FindOneByID(ctx context.Context, userID uint) (*models.User, error)
	FindOneByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateOne(ctx context.Context, user *models.User) error
}
