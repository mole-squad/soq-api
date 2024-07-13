package agendas

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mole-squad/soq-api/pkg/db"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type AgendaRepoParams struct {
	fx.In

	DBService     interfaces.DBService
	LoggerService interfaces.LoggerService
}

type AgendaRepoResult struct {
	fx.Out

	AgendaRepo interfaces.AgendaRepo
}

type AgendaRepo struct {
	dbService interfaces.DBService
	logger    interfaces.LoggerService
}

func NewAgendaRepo(params AgendaRepoParams) (AgendaRepoResult, error) {
	repo := &AgendaRepo{
		dbService: params.DBService,
		logger:    params.LoggerService,
	}

	return AgendaRepoResult{AgendaRepo: repo}, nil
}

func (repo *AgendaRepo) CreateOne(ctx context.Context, agenda *models.Agenda) error {
	err := repo.dbService.CreateOne(ctx, agenda)
	if err != nil {
		return fmt.Errorf("failed to create one agenda: %w", err)
	}

	repo.logger.Debug("Created one agenda", "agenda", agenda)

	return nil
}

func (repo *AgendaRepo) UpdateOne(ctx context.Context, agenda *models.Agenda) error {
	err := repo.dbService.UpdateOne(ctx, agenda)
	if err != nil {
		return fmt.Errorf("failed to update one agenda: %w", err)
	}

	repo.logger.Debug("Updated one agenda", "agenda", agenda)

	return nil
}

func (repo *AgendaRepo) FindManyByUser(ctx context.Context, userID uint) ([]models.Agenda, error) {
	var agendas []models.Agenda

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
