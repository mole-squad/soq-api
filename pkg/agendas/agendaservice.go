package agendas

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/burkel24/go-mochi"

	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

const (
	// Max time into the future to generate agendas for,
	// relative to the TimeWindow start time
	TimeWindowPreGenerationTimeMins = 5
)

type AgendaServiceParams struct {
	fx.In

	AgendaRepo          interfaces.AgendaRepo
	FocusAreaService    interfaces.FocusAreaService
	LoggerService       mochi.LoggerService
	NotificationService interfaces.NotificationService
	QuotaService        interfaces.QuotaService
	TaskService         interfaces.TaskService
	UserService         mochi.UserService
}

type AgendaServiceResult struct {
	fx.Out

	AgendaService interfaces.AgendaService
}

type AgendaService struct {
	mochi.Service[*models.Agenda]

	agendaRepo          interfaces.AgendaRepo
	focusAreaService    interfaces.FocusAreaService
	logger              mochi.LoggerService
	notificationService interfaces.NotificationService
	quotaService        interfaces.QuotaService
	taskService         interfaces.TaskService
	userService         mochi.UserService
}

func NewAgendaService(params AgendaServiceParams) (AgendaServiceResult, error) {
	embeddedSvc := mochi.NewService(
		params.AgendaRepo,
	)

	srv := &AgendaService{
		Service: embeddedSvc,

		focusAreaService:    params.FocusAreaService,
		logger:              params.LoggerService,
		notificationService: params.NotificationService,
		quotaService:        params.QuotaService,
		taskService:         params.TaskService,
		userService:         params.UserService,
	}

	return AgendaServiceResult{AgendaService: srv}, nil
}

func (srv *AgendaService) GenerateAgendasForUpcomingTimeWindows(ctx context.Context) error {
	srv.logger.Info("Generating agendas for upcoming time windows")

	users, err := srv.userService.ListUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to load users: %w", err)
	}

	for _, user := range users {
		err := srv.GenerateAgendasForUser(ctx, user.(*models.User))

		if err != nil {
			return fmt.Errorf("failed to generate agendas for user: %w", err)
		}
	}

	return nil
}

func (srv *AgendaService) GenerateAgendasForUser(ctx context.Context, user *models.User) error {
	srv.logger.Info("Generating agendas for user", "user", user.ID)

	focusAreas, err := srv.focusAreaService.ListByUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to load focus areas for user: %w", err)
	}

	for _, focusArea := range focusAreas {
		err := srv.GenerateAgendasForFocusArea(ctx, user, focusArea)
		if err != nil {
			return fmt.Errorf("failed to generate agendas for focus area: %w", err)
		}
	}

	return nil
}

func (srv *AgendaService) GenerateAgendasForFocusArea(ctx context.Context, user *models.User, focusArea *models.FocusArea) error {
	srv.logger.Info("Generating agendas for focus area", "user", user.ID, "focusArea", focusArea.ID)

	location, err := time.LoadLocation(user.Timezone)
	if err != nil {
		return fmt.Errorf("failed to load user timezone: %w", err)
	}

	timeWindow := srv.getTimeWindowForGenerationWindow(focusArea, location)
	if timeWindow == nil {
		srv.logger.Debug("Focus area is not within generation window", "focusArea", focusArea.ID)

		return nil
	}

	timeRangeStart, timeRangeEnd := srv.getNewAgendaTimeRange(timeWindow, location)

	existingAgenda, err := srv.agendaRepo.FindOneByTimeRangeFocusArea(ctx, user.ID, focusArea.ID, timeRangeStart, timeRangeEnd)
	if err != nil {
		return fmt.Errorf("failed to load existing agenda: %w", err)
	}

	if existingAgenda != nil {
		srv.logger.Info(
			"Agenda already exists for time range",
			"agenda", existingAgenda.ID,
			"items", len(existingAgenda.AgendaItems),
			"focusArea", focusArea.ID,
			"startTime", timeRangeStart,
			"endTime", timeRangeEnd,
			"user", user.ID,
		)

		return nil
	}

	agenda := &models.Agenda{
		UserID:      user.ID,
		FocusAreaID: focusArea.ID,
		Status:      models.AgendaStatusPending,
		StartTime:   timeRangeStart,
		EndTime:     timeRangeEnd,
	}

	err = srv.agendaRepo.CreateOne(ctx, agenda)
	if err != nil {
		return fmt.Errorf("failed to create agenda: %w", err)
	}

	return nil
}

func (srv *AgendaService) PopulatePendingAgendas(ctx context.Context) error {
	pendingAgendas, err := srv.agendaRepo.FindManyByStatus(ctx, models.AgendaStatusPending)
	if err != nil {
		return fmt.Errorf("failed to load pending agendas: %w", err)
	}

	for _, agenda := range pendingAgendas {
		err := srv.populateAgenda(ctx, &agenda)
		if err != nil {
			return fmt.Errorf("failed to populate agenda: %w", err)
		}
	}

	return nil
}

func (srv *AgendaService) SendAgendaNotifications(ctx context.Context) error {
	agendas, err := srv.agendaRepo.FindManyByStatus(ctx, models.AgendaStatusGenerated)
	if err != nil {
		return fmt.Errorf("failed to load generated agendas: %w", err)
	}

	if len(agendas) == 0 {
		srv.logger.Debug("No agendas to send")
	}

	sendErrors := make([]error, 0)

	for _, agenda := range agendas {
		err := srv.sendAgendaNotification(ctx, &agenda)
		if err != nil {
			sendErrors = append(sendErrors, fmt.Errorf("failed to send agenda notification: %w", err))
		}
	}

	if len(sendErrors) > 0 {
		return errors.Join(sendErrors...)
	}

	return nil
}

func (srv *AgendaService) getTimeWindowForGenerationWindow(focusArea *models.FocusArea, tz *time.Location) *models.TimeWindow {
	for _, timeWindow := range focusArea.TimeWindows {
		if srv.isTimeWindowWithinGenerationWindow(&timeWindow, tz) {
			srv.logger.Debug(
				"Time window is within generation window",
				"startTime", timeWindow.StartTime,
				"endTime", timeWindow.EndTime,
				"weekdays", timeWindow.Weekdays,
			)

			return &timeWindow
		}
	}

	return nil
}

func (srv *AgendaService) isTimeWindowWithinGenerationWindow(timeWindow *models.TimeWindow, tz *time.Location) bool {
	now := time.Now().In(tz)
	today := now.Weekday()

	hasMatchingDay := false

	for _, day := range timeWindow.Weekdays {
		if day == int32(today) {
			hasMatchingDay = true
			break
		}
	}

	if !hasMatchingDay {
		return false
	}

	nowFloat := float32(now.Hour()) + float32(now.Minute())/60

	if nowFloat > timeWindow.EndTime {
		return false
	}

	return nowFloat >= timeWindow.StartTime-(TimeWindowPreGenerationTimeMins/60)
}

func (srv *AgendaService) getNewAgendaTimeRange(window *models.TimeWindow, tz *time.Location) (time.Time, time.Time) {
	now := time.Now().In(tz)

	timeRangeStart := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		int(window.StartTime),
		0,
		0,
		0,
		tz,
	)

	timeRangeEnd := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		int(window.EndTime),
		0,
		0,
		0,
		tz,
	)

	return timeRangeStart, timeRangeEnd
}

func (srv *AgendaService) populateAgenda(ctx context.Context, agenda *models.Agenda) error {
	srv.logger.Info(
		"Populating agenda",
		"agenda",
		agenda.ID,
		"focusArea",
		agenda.FocusAreaID,
		"user",
		agenda.UserID,
	)

	tasks, err := srv.taskService.ListOpenUserTasksForFocusArea(ctx, agenda.User.ID, agenda.FocusAreaID)
	if err != nil {
		return fmt.Errorf("failed to load open tasks: %w", err)
	}

	var agendaItems []models.AgendaItem

	for _, task := range tasks {
		agendaItem := models.AgendaItem{
			AgendaID: agenda.ID,
			TaskID:   &task.ID,
		}

		agendaItems = append(agendaItems, agendaItem)
	}

	agenda.AgendaItems = agendaItems
	agenda.Status = models.AgendaStatusGenerated

	err = srv.agendaRepo.UpdateOne(ctx, agenda.UserID, agenda)
	if err != nil {
		return fmt.Errorf("failed to update agenda: %w", err)
	}

	return nil
}

func (srv *AgendaService) sendAgendaNotification(ctx context.Context, agenda *models.Agenda) error {
	srv.logger.Info(
		"Sending agenda notification",
		"agenda",
		agenda.ID,
		"focusArea",
		agenda.FocusAreaID,
		"user",
		agenda.UserID,
		"startTime",
		agenda.StartTime,
		"endTime",
		agenda.EndTime,
	)

	if len(agenda.AgendaItems) == 0 {
		srv.logger.Info("Agenda has no items", "agenda", agenda.ID)
		return nil
	}

	err := srv.notificationService.SendNotification(
		ctx,
		agenda.UserID,
		agenda.GetTitle(),
		agenda.GetBody(),
	)

	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	agenda.Status = models.AgendaStatusSent

	err = srv.agendaRepo.UpdateOne(ctx, agenda.UserID, agenda)
	if err != nil {
		return fmt.Errorf("failed to update agenda: %w", err)
	}

	return nil
}
