package generics

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
)

type Repository[M models.Model] struct {
	db     interfaces.DBService
	logger interfaces.LoggerService

	joinTables    []string
	preloadTables []string
	tableName     string
}

type RepositoryOption[M models.Model] func(*Repository[M])

func NewRepository[M models.Model](
	db interfaces.DBService,
	logger interfaces.LoggerService,
	opts ...RepositoryOption[M],
) interfaces.Repository[M] {
	repo := &Repository[M]{
		db:     db,
		logger: logger,
	}

	for _, opt := range opts {
		opt(repo)
	}

	return repo
}

func (r *Repository[M]) FindOne(ctx context.Context, query string, args ...interface{}) (M, error) {
	var item M

	err := r.db.FindOne(ctx, &item, r.joinTables, []string{}, query, args...)
	if err != nil {
		return item, fmt.Errorf("failed to find one item: %w", err)
	}

	r.logger.Debug("Found one item", "item", item.GetID(), "table", r.tableName)

	return item, nil
}

func (r *Repository[M]) FindOneByID(ctx context.Context, itemID uint, query string, args ...interface{}) (M, error) {
	fullQuery := fmt.Sprintf("%s.id = ?", r.tableName)
	if query != "" {
		fullQuery = fmt.Sprintf("%s AND %s", fullQuery, query)
	}

	fullArgs := append([]interface{}{itemID}, args...)

	return r.FindOne(ctx, fullQuery, fullArgs...)
}

func (r *Repository[M]) FindOneByUser(ctx context.Context, userID uint, query string, args ...interface{}) (M, error) {
	var item M

	fullQuery := fmt.Sprintf("%s.user_id = ?", r.tableName)
	if query != "" {
		fullQuery = fmt.Sprintf("%s AND %s", fullQuery, query)
	}

	fullArgs := append([]interface{}{userID}, args...)

	err := r.db.FindOne(ctx, &item, r.joinTables, []string{}, fullQuery, fullArgs...)
	if err != nil {
		return item, fmt.Errorf("failed to find one item: %w", err)
	}

	r.logger.Debug("Found one item by user", "item", item.GetID(), "table", r.tableName)

	return item, nil
}

func (r *Repository[M]) FindManyByUser(ctx context.Context, userID uint, query string, args ...interface{}) ([]M, error) {
	var items []M

	fullQuery := fmt.Sprintf("%s.user_id = ?", r.tableName)
	if query != "" {
		fullQuery = fmt.Sprintf("%s AND %s", fullQuery, query)
	}

	fullArgs := append([]interface{}{userID}, args...)

	err := r.db.FindMany(ctx, &items, r.joinTables, r.preloadTables, fullQuery, fullArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to find many items by user: %w", err)
	}

	r.logger.Debug("Found many items by user", "table", r.tableName, "count", len(items))

	return items, nil
}

func (r *Repository[M]) CreateOne(ctx context.Context, item M) error {
	err := r.db.CreateOne(ctx, item)
	if err != nil {
		return fmt.Errorf("failed to create one item: %w", err)
	}

	r.logger.Debug("Created one item", "item", item.GetID(), "table", r.tableName)

	return nil
}

func (r *Repository[M]) UpdateOne(ctx context.Context, itemID uint, item M) error {
	err := r.db.UpdateOne(ctx, itemID, item)
	if err != nil {
		return fmt.Errorf("failed to update one item: %w", err)
	}

	r.logger.Debug("Updated one item", "item", item.GetID(), "table", r.tableName)

	return nil
}

func (r *Repository[M]) DeleteOne(ctx context.Context, itemID uint) error {
	item := new(M)

	err := r.db.DeleteOne(ctx, itemID, item)
	if err != nil {
		return fmt.Errorf("failed to delete one item: %w", err)
	}

	r.logger.Debug("Deleted one item", "item", itemID, "table", r.tableName)

	return nil
}

func WithTableName[M models.Model](tableName string) RepositoryOption[M] {
	return func(r *Repository[M]) {
		r.tableName = tableName
	}
}

func WithJoinTables[M models.Model](joinTables ...string) RepositoryOption[M] {
	return func(r *Repository[M]) {
		r.joinTables = joinTables
	}
}

func WithPreloadTables[M models.Model](preloadTables ...string) RepositoryOption[M] {
	return func(r *Repository[M]) {
		r.preloadTables = preloadTables
	}
}
