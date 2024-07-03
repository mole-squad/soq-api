package auth

import (
	"context"

	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/users"
	"gorm.io/gorm"
)

func GetUserFromCtx(ctx context.Context) (interfaces.User, error) {
	return &users.User{
		Model: gorm.Model{ID: 1},
	}, nil
}
