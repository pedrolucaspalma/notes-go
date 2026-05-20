package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func NewLogger() (*AppLogger, error) {
	debugFile, err := os.OpenFile("go-notes-debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		return nil, fmt.Errorf("opening debug file: %w", err)
	}

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logHandler := slog.NewJSONHandler(debugFile, opts)

	logger := slog.New(logHandler)

	appLogger := AppLogger{
		slogLogger: logger,
	}
	return &appLogger, nil
}

type AppLogger struct {
	slogLogger *slog.Logger
}

func (l *AppLogger) Debug(ctx context.Context, msg string, args ...any) {
	l.slogLogger.DebugContext(ctx, msg, args...)
}

func (l *AppLogger) Info(ctx context.Context, msg string, args ...any) {
	l.slogLogger.InfoContext(ctx, msg, args...)
}

func (l *AppLogger) Warn(ctx context.Context, msg string, args ...any) {
	l.slogLogger.WarnContext(ctx, msg, args...)
}

func (l *AppLogger) Error(ctx context.Context, msg string, args ...any) {
	l.slogLogger.ErrorContext(ctx, msg, args...)
}
