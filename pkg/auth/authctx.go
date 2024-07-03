package auth

import (
	"context"

	"github.com/burkel24/task-app/pkg/models"
	"gorm.io/gorm"
)

func GetUserFromCtx(ctx context.Context) (models.User, error) {
	return models.User{
		Model: gorm.Model{ID: 1},
	}, nil
}
