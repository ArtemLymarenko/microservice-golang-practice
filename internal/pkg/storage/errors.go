package storage

import "errors"

var (
	ErrHealthCheckPostgres     = errors.New("error health check postgres")
	ErrCreateMigrationPostgres = errors.New("error creating a migration postgres")
	ErrMigratePostgres         = errors.New("error migrating postgres")
	ErrGetPostgresConnection   = errors.New("error getting postgres connection")
	ErrClosePostgresConnection = errors.New("error closing postgres connection")
)
