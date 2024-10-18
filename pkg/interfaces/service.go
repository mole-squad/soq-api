package interfaces

import "context"

type Service[M Resource] interface {
	ListByUser(ctx context.Context, userID uint) ([]M, error)
	CreateOne(ctx context.Context, userID uint, item M) (M, error)
	GetOne(ctx context.Context, itemID uint) (M, error)
	UpdateOne(ctx context.Context, itemID uint, item M) (M, error)
	DeleteOne(ctx context.Context, itemID uint) error
}
