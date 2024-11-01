package postgres

import "errors"

var (
	ErrHealthCheckPostgres     = errors.New("error health check postgres")
	ErrCreateMigrationPostgres = errors.New("error creating a migration postgres")
	ErrMigratePostgres         = errors.New("error migrating postgres")
	ErrGetConnection           = errors.New("error getting postgres connection")
	ErrCloseConnection         = errors.New("error closing postgres connection")
)
