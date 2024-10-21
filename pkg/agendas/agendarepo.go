package agendas

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/burkel24/go-mochi"

	"github.com/mole-squad/soq-api/pkg/db"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type AgendaRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService mochi.LoggerService
}

type AgendaRepoResult struct {
	fx.Out

	AgendaRepo interfaces.AgendaRepo
}

type AgendaRepo struct {
	mochi.Repository[*models.Agenda]

	dbService interfaces.DBService
	logger    mochi.LoggerService
}

func NewAgendaRepo(params AgendaRepoParams) (AgendaRepoResult, error) {
	embeddedRepo := mochi.NewRepository(
		params.DBService,
		params.LoggerService,
		mochi.WithTableName[*models.Agenda]("agendas"),
	)

	repo := &AgendaRepo{
		Repository: embeddedRepo,
	}

	return AgendaRepoResult{AgendaRepo: repo}, nil
}

func (repo *AgendaRepo) FindManyByUser(ctx context.Context, userID uint, query string, args ...interface{}) ([]*models.Agenda, error) {
	var agendas []*models.Agenda

	err := repo.dbService.FindMany(ctx, &agendas, []string{"FocusArea"}, []string{}, "agendas.user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list agendas: %w", err)
	}

	return agendas, nil
}

func (repo *AgendaRepo) FindManyByStatus(ctx context.Context, status models.AgendaStatus) ([]models.Agenda, error) {
	var agendas []models.Agenda

	err := repo.dbService.FindMany(
		ctx,
		&agendas,
		[]string{"FocusArea", "User"},
		[]string{"AgendaItems", "AgendaItems.Task", "AgendaItems.Quota"},
		"agendas.status = ?",
		status,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to list pending agendas: %w", err)
	}

	return agendas, nil

}

func (repo *AgendaRepo) FindOneByTimeRangeFocusArea(
	ctx context.Context,
	userID uint,
	focusAreaID uint,
	startTime, endTime time.Time,
) (*models.Agenda, error) {
	var agenda models.Agenda

	err := repo.dbService.FindOne(
		ctx,
		&agenda,
		[]string{},
		[]string{"AgendaItems"},
		"agendas.user_id = ? AND agendas.focus_area_id = ? AND agendas.start_time = ? AND agendas.end_time = ?",
		userID,
		focusAreaID,
		startTime,
		endTime,
	)

	if err != nil {
		if errors.Is(err, db.ErrorNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find agenda: %w", err)
	}

	return &agenda, nil
}
