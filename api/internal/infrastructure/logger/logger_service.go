package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"task-management/internal/application/service"
	"time"
)

type logger struct {
	logger *slog.Logger
}

func NewLogger() service.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: replacer,
	})
	l := slog.New(jsonHandler)
	slog.SetDefault(l)
	return &logger{
		logger: l,
	}
}

func replacer(groups []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.LevelKey:
		a.Key = "severity"
		if level := a.Value.Any().(slog.Level); level == slog.LevelWarn {
			a.Value = slog.StringValue("WARNING")
		}
	case slog.TimeKey:
		a.Key = "timestamp"
	case slog.MessageKey:
		a.Key = "message"
	case slog.SourceKey:
		a.Key = "logging.googleapis.com/sourceLocation"
	}

	return a
}

func (l *logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.logger.Enabled(ctx, level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])

	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)

	_ = l.logger.Handler().Handle(ctx, r)
}

func (l *logger) Info(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelInfo, msg, args...)
}

func (l *logger) Infof(ctx context.Context, format string, args ...any) {
	l.log(ctx, slog.LevelInfo, fmt.Sprintf(format, args...))
}

func (l *logger) Error(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelError, msg, args...)
}

func (l *logger) Errorf(ctx context.Context, format string, args ...any) {
	l.log(ctx, slog.LevelError, fmt.Sprintf(format, args...))
}
