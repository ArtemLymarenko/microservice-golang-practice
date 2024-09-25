package storage

import "errors"

var (
	ErrGetPostgresConnection   = errors.New("error getting postgres connection")
	ErrClosePostgresConnection = errors.New("error closing postgres connection")
)
