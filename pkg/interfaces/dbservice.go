package interfaces

import "context"

type DBService interface {
	FindMany(ctx context.Context, result interface{}, query interface{}, args ...interface{}) error
}
