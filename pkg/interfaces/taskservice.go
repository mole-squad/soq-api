package interfaces

import "context"

type TaskService interface {
	ListUserTasks(ctx context.Context, user User) ([]Task, error)
}
