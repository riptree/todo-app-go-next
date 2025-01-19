package db

import (
	"context"
	"database/sql"
	"task-management/internal/application/service"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

func NewTransaction(conn *bun.DB) service.Transaction {
	return &transaction{
		conn: conn,
	}
}

type transaction struct {
	conn *bun.DB
}

func (r *transaction) WithinTransaction(ctx context.Context, f func(ctx context.Context) error) (err error) {
	tx, err := GetTxOrDB(ctx, r.conn).BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("panic error %v", r)
			tx.Rollback()
		}
	}()

	err = f(SetTx(ctx, &tx))
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

type dbTx struct{}

func SetTx(ctx context.Context, tx *bun.Tx) context.Context {
	return context.WithValue(ctx, dbTx{}, tx)
}

func getTx(ctx context.Context) *bun.Tx {
	if tx, ok := ctx.Value(dbTx{}).(*bun.Tx); ok {
		return tx
	}
	return nil
}

func GetTxOrDB(ctx context.Context, db *bun.DB) bun.IDB {
	if tx := getTx(ctx); tx != nil {
		return tx
	}
	return db
}
