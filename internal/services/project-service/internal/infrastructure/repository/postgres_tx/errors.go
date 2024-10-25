package postgresTx

import "errors"

var (
	ErrCommitTrx = errors.New("failed to finish transaction")
)
