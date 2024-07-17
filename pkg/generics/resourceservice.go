package generics

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
)

type ResourceService[M interfaces.Resource] struct {
	repo interfaces.ResourceRepository[M]

	listQuery *Query
	getQuery  *Query
}

type ResourceServiceOption[M interfaces.Resource] func(*ResourceService[M])

func NewResourceService[M interfaces.Resource](
	repo interfaces.ResourceRepository[M],
	opts ...ResourceServiceOption[M],
) interfaces.ResourceService[M] {
	svc := &ResourceService[M]{
		repo: repo,
	}

	for _, opt := range opts {
		opt(svc)
	}

	if svc.listQuery == nil {
		svc.listQuery = &Query{}
	}

	if svc.getQuery == nil {
		svc.getQuery = &Query{}
	}

	return svc
}

func (s *ResourceService[M]) ListByUser(ctx context.Context, userID uint) ([]M, error) {
	items, err := s.repo.FindManyByUser(ctx, userID, s.listQuery.Filter, s.listQuery.Args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list user items: %w", err)
	}

	return items, nil
}

func (s *ResourceService[M]) CreateOne(ctx context.Context, userID uint, item M) (M, error) {
	item.SetUserID(userID)

	err := s.repo.CreateOne(ctx, item)
	if err != nil {
		return item, fmt.Errorf("failed to create user task: %w", err)
	}

	return item, nil
}

func (s *ResourceService[M]) GetOne(ctx context.Context, itemID uint) (M, error) {
	item, err := s.repo.FindOneByID(ctx, itemID, s.getQuery.Filter, s.getQuery.Args...)
	if err != nil {
		return item, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

func (s *ResourceService[M]) UpdateOne(ctx context.Context, itemID uint, item M) (M, error) {
	err := s.repo.UpdateOne(ctx, itemID, item)
	if err != nil {
		return item, fmt.Errorf("failed to update user task: %w", err)
	}

	return item, nil
}

func (s *ResourceService[M]) DeleteOne(ctx context.Context, itemID uint) error {
	err := s.repo.DeleteOne(ctx, itemID)
	if err != nil {
		return fmt.Errorf("failed to delete user task: %w", err)
	}

	return nil
}

func WithListQuery[M interfaces.Resource](query string, args ...interface{}) ResourceServiceOption[M] {
	return func(s *ResourceService[M]) {
		s.listQuery = &Query{
			Filter: query,
			Args:   args,
		}
	}
}

func WithGetQuery[M interfaces.Resource](query string, args ...interface{}) ResourceServiceOption[M] {
	return func(s *ResourceService[M]) {
		s.getQuery = &Query{
			Filter: query,
			Args:   args,
		}
	}
}
