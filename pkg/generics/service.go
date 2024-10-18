package generics

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
)

type Service[M interfaces.Resource] struct {
	repo interfaces.Repository[M]

	listQuery *Query
	getQuery  *Query
}

type ServiceOption[M interfaces.Resource] func(*Service[M])

func NewService[M interfaces.Resource](
	repo interfaces.Repository[M],
	opts ...ServiceOption[M],
) interfaces.Service[M] {
	svc := &Service[M]{
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

func (s *Service[M]) ListByUser(ctx context.Context, userID uint) ([]M, error) {
	items, err := s.repo.FindManyByUser(ctx, userID, s.listQuery.Filter, s.listQuery.Args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list user items: %w", err)
	}

	return items, nil
}

func (s *Service[M]) CreateOne(ctx context.Context, userID uint, item M) (M, error) {
	item.SetUserID(userID)

	err := s.repo.CreateOne(ctx, item)
	if err != nil {
		return item, fmt.Errorf("failed to create user task: %w", err)
	}

	return item, nil
}

func (s *Service[M]) GetOne(ctx context.Context, itemID uint) (M, error) {
	item, err := s.repo.FindOneByID(ctx, itemID, s.getQuery.Filter, s.getQuery.Args...)
	if err != nil {
		return item, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

func (s *Service[M]) UpdateOne(ctx context.Context, itemID uint, item M) (M, error) {
	err := s.repo.UpdateOne(ctx, itemID, item)
	if err != nil {
		return item, fmt.Errorf("failed to update user task: %w", err)
	}

	return item, nil
}

func (s *Service[M]) DeleteOne(ctx context.Context, itemID uint) error {
	err := s.repo.DeleteOne(ctx, itemID)
	if err != nil {
		return fmt.Errorf("failed to delete user task: %w", err)
	}

	return nil
}

func WithListQuery[M interfaces.Resource](query string, args ...interface{}) ServiceOption[M] {
	return func(s *Service[M]) {
		s.listQuery = &Query{
			Filter: query,
			Args:   args,
		}
	}
}

func WithGetQuery[M interfaces.Resource](query string, args ...interface{}) ServiceOption[M] {
	return func(s *Service[M]) {
		s.getQuery = &Query{
			Filter: query,
			Args:   args,
		}
	}
}
