package agendas

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/burkel24/task-app/pkg/db"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
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

func (repo *AgendaRepo) FindManyByUser(ctx context.Context, userID uint) ([]models.Agenda, error) {
	var agendas []models.Agenda

	err := repo.dbService.FindMany(ctx, &agendas, []string{"FocusArea"}, []string{}, "agendas.user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list agendas: %w", err)
	}

	return agendas, nil
}

func (repo *AgendaRepo) FindManyByPending(ctx context.Context) ([]models.Agenda, error) {
	var agendas []models.Agenda

	err := repo.dbService.FindMany(ctx, &agendas, []string{"FocusArea", "User"}, []string{}, "agendas.status = ?", models.AgendaStatusPending)
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
		[]string{},
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
