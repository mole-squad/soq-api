package interfaces

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
)

type UserService interface {
	ListUsers(ctx context.Context) ([]models.User, error)
}
