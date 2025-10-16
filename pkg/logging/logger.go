package logging

import (
	"log/slog"
)

type Level = slog.Level

const (
	DebugLevel  = slog.LevelDebug
	InfoLevel   = slog.LevelInfo
	WarnLevel   = slog.LevelWarn
	ErrorLevel  = slog.LevelError
	DPanicLevel = slog.LevelError
	PanicLevel  = slog.LevelError
	FatalLevel  = slog.LevelError
)

// Logger is used for logging formatted messages.
type Logger interface {
	// Debugf logs messages at DEBUG level.
	Debugf(format string, args ...any)
	// Infof logs messages at INFO level.
	Infof(format string, args ...any)
	// Warnf logs messages at WARN level.
	Warnf(format string, args ...any)
	// Errorf logs messages at ERROR level.
	Errorf(format string, args ...any)
	// Fatalf logs messages at FATAL level.
	Fatalf(format string, args ...any)
}
