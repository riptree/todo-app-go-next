package service

import (
	"context"
)

type Logger interface {
	Info(ctx context.Context, msg string, args ...any)
	Infof(ctx context.Context, format string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Errorf(ctx context.Context, format string, args ...any)
}
