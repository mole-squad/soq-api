package interfaces

import "context"

type DBService interface {
	CreateOne(ctx context.Context, record interface{}) error
	UpdateOne(ctx context.Context, record interface{}) error
	DeleteOne(ctx context.Context, record interface{}) error
	FindMany(ctx context.Context, result interface{}, joins []string, query interface{}, args ...interface{}) error
}
