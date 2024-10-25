package postgres

import "errors"

var (
	ErrFinishTx = errors.New("failed to finish transaction")
)
