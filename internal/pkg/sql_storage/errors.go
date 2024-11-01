package sqlStorage

import "errors"

var (
	ErrFinishTx    = errors.New("failed to finish transaction")
	ErrRowsNotRead = errors.New("rows were not read")
)
