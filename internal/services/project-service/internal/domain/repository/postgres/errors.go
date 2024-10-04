package postgres

import "errors"

var (
	ErrQueryRowWithContext = errors.New("error querying row with context")
	ErrExecWithContext     = errors.New("error exec with context")
	ErrCommitTrx           = errors.New("error commiting transaction")
)
