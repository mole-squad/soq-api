package notifications

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gregdel/pushover"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
)

type NotificationServiceParams struct {
	fx.In

	DeviceRepo    interfaces.DeviceRepo
	LoggerService interfaces.LoggerService
}

type NotificationServiceResult struct {
	fx.Out

	NotificationService interfaces.NotificationService
}

type NotificationService struct {
	client               *pushover.Pushover
	deviceRepo           interfaces.DeviceRepo
	logger               interfaces.LoggerService
	notificationsEnabled bool
}

func NewNotificationService(p NotificationServiceParams) (NotificationServiceResult, error) {
	var result NotificationServiceResult

	token := os.Getenv("PUSHOVER_TOKEN")
	if token == "" {
		return result, fmt.Errorf("PUSHOVER_TOKEN env var is required")
	}

	notificationsEnabled := os.Getenv("NOTIFICATIONS_ENABLED") == "true"

	slog.Info("Creating pushover client", "token", token)
	client := pushover.New(token)

	result.NotificationService = &NotificationService{
		client:               client,
		deviceRepo:           p.DeviceRepo,
		logger:               p.LoggerService,
		notificationsEnabled: notificationsEnabled,
	}

	return result, nil
}

func (srv *NotificationService) SendNotification(
	ctx context.Context,
	userID uint,
	title string,
	message string,
) error {
	srv.logger.Info(
		"Sending notification",
		"userID",
		userID,
		"title",
		title,
		"message",
		message,
		"noficiationsEnabled",
		srv.notificationsEnabled,
	)

	userDevices, err := srv.deviceRepo.FindManyByUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to send notifications: %w", err)
	}

	for _, device := range userDevices {
		err := srv.sendNotificationToDevice(&device, title, message)
		if err != nil {
			return fmt.Errorf("failed to send notifications: %w", err)
		}
	}

	return nil
}

func (srv *NotificationService) sendNotificationToDevice(device *models.Device, title string, message string) error {
	srv.logger.Info(
		"Sending notification to device",
		"title",
		title,
		"message",
		message,
		"device",
		device.DeviceID,
		"notificationsEnabled",
		srv.notificationsEnabled,
	)

	if !srv.notificationsEnabled {
		return nil
	}

	recipient := pushover.NewRecipient(device.UserKey)

	msg := pushover.Message{
		Message:    message,
		Title:      title,
		DeviceName: device.DeviceID,
	}

	_, err := srv.client.SendMessage(&msg, recipient)
	if err != nil {
		return fmt.Errorf("failed to send notification to pushover: %w", err)
	}

	return nil
}
