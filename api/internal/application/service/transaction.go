package service

import (
	"context"
)

type Transaction interface {
	WithinTransaction(ctx context.Context, f func(ctx context.Context) error) (err error)
}
