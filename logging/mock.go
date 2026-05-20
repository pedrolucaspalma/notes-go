package logging

import "context"

type MockLogger struct{}

func (l *MockLogger) Debug(ctx context.Context, msg string, args ...any) {}
func (l *MockLogger) Info(ctx context.Context, msg string, args ...any)  {}
func (l *MockLogger) Warn(ctx context.Context, msg string, args ...any)  {}
func (l *MockLogger) Error(ctx context.Context, msg string, args ...any) {}
