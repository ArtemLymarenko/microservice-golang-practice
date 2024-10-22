package postgres

import "errors"

var (
	ErrQueryRowWithContext = errors.New("failed to query project data")
	ErrProjectNotFound     = errors.New("project was not found")
	ErrCommitTrx           = errors.New("failed to finish transaction")
)
