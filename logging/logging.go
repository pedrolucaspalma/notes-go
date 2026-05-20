package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func NewLogger() (*AppLogger, error) {
	debugFile, err := os.OpenFile("go-notes-debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		return nil, fmt.Errorf("opening debug file: %w", err)
	}
	writter := formattedJSONFileWritter{
		destination: debugFile,
	}

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logHandler := slog.NewJSONHandler(&writter, opts)

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

type formattedJSONFileWritter struct {
	destination io.Writer
}

func (t *formattedJSONFileWritter) Write(p []byte) (int, error) {
	var obj any
	err := json.Unmarshal(p, &obj)
	if err != nil {
		return 0, err
	}

	identedJson, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return 0, err
	}

	identedJson = append(identedJson, '\n')
	return t.destination.Write(identedJson)
}
