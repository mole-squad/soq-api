package interfaces

import "log/slog"

type LoggerService interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Logger() *slog.Logger
}
