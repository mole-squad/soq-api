package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type UserService interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, userID uint) (*models.User, error)
	GetUserByCredentials(ctx context.Context, username, passwordHash string) (*models.User, error)
}
