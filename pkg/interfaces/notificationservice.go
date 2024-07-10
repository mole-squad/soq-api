package interfaces

import "context"

type NotificationService interface {
	SendNotification(
		ctx context.Context,
		userID uint,
		title string,
		message string,
	) error
}
