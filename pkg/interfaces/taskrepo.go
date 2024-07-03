package interfaces

import "context"

type TaskRepo interface {
	FindManyTasksByUser(ctx context.Context, userID uint) ([]Task, error)
}
