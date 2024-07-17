package db

import (
	"context"
	"fmt"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
)

type Repo[M models.Model] struct {
	db     interfaces.DBService
	logger interfaces.LoggerService

	joinTables []string
	tableName  string
}

type RepoOption[M models.Model] func(*Repo[M])

func NewRepo[M models.Model](
	db interfaces.DBService,
	logger interfaces.LoggerService,
	opts ...RepoOption[M],
) *Repo[M] {
	repo := &Repo[M]{
		db:     db,
		logger: logger,
	}

	for _, opt := range opts {
		opt(repo)
	}

	return repo
}

func (r *Repo[M]) FindOneByUser(ctx context.Context, userID uint, query string, args ...interface{}) (M, error) {
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

func (r *Repo[M]) FindManyByUser(ctx context.Context, userID uint, query string, args ...interface{}) ([]M, error) {
	var items []M

	fullQuery := fmt.Sprintf("%s.user_id = ?", r.tableName)
	if query != "" {
		fullQuery = fmt.Sprintf("%s AND %s", fullQuery, query)
	}

	fullArgs := append([]interface{}{userID}, args...)

	err := r.db.FindMany(ctx, &items, r.joinTables, []string{}, fullQuery, fullArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to find many items by user: %w", err)
	}

	r.logger.Debug("Found many items by user", "table", r.tableName, "count", len(items))

	return items, nil
}

func (r *Repo[M]) CreateOne(ctx context.Context, item M) error {
	err := r.db.CreateOne(ctx, item)
	if err != nil {
		return fmt.Errorf("failed to create one item: %w", err)
	}

	r.logger.Debug("Created one item", "item", item.GetID(), "table", r.tableName)

	return nil
}

func (r *Repo[M]) UpdateOne(ctx context.Context, item M) error {
	err := r.db.UpdateOne(ctx, item)
	if err != nil {
		return fmt.Errorf("failed to update one item: %w", err)
	}

	r.logger.Debug("Updated one item", "item", item.GetID(), "table", r.tableName)

	return nil
}

func (r *Repo[M]) DeleteOne(ctx context.Context, itemID uint) error {
	item := new(M)

	err := r.db.DeleteOne(ctx, itemID, item)
	if err != nil {
		return fmt.Errorf("failed to delete one item: %w", err)
	}

	r.logger.Debug("Deleted one item", "item", itemID, "table", r.tableName)

	return nil
}

func WithTableName[M models.Model](tableName string) RepoOption[M] {
	return func(r *Repo[M]) {
		r.tableName = tableName
	}
}

func WithJoinTables[M models.Model](joinTables ...string) RepoOption[M] {
	return func(r *Repo[M]) {
		r.joinTables = joinTables
	}
}
