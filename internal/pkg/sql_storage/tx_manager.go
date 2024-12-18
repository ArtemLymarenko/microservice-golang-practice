package sqlStorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type TxManager interface {
	Run(
		ctx context.Context,
		callback func(ctx context.Context, tx *sql.Tx) error,
	) error
}

type SQLTransactionManager struct {
	db *sql.DB
}

func NewTxManager(db *sql.DB) *SQLTransactionManager {
	return &SQLTransactionManager{db: db}
}

func (m *SQLTransactionManager) Run(
	ctx context.Context,
	runTx func(ctx context.Context, tx *sql.Tx) error,
) (rErr error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.Join(ErrFinishTx)
	}

	defer func() {
		if rErr != nil {
			rErr = errors.Join(ErrFinishTx, tx.Rollback())
		}
	}()

	defer func() {
		if rec := recover(); rec != nil {
			if e, ok := rec.(error); ok {
				rErr = e
			} else {
				rErr = fmt.Errorf("%s", rec)
			}
		}
	}()

	if err = runTx(ctx, tx); err != nil {
		return err
	}

	return errors.Join(ErrFinishTx, tx.Commit())
}
