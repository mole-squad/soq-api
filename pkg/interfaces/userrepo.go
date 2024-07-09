package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type UserRepo interface {
	ListUsers(ctx context.Context) ([]models.User, error)
}
