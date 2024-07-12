package auth

import (
	"context"
	"fmt"

	"github.com/burkel24/task-app/pkg/models"
)

func GetUserFromCtx(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		return nil, fmt.Errorf("could not get user from context")
	}

	return user, nil
}
