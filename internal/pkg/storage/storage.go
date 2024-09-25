package storage

import "database/sql"

type Postgres interface {
	CloseConnection() error
	GetConnection() (*sql.DB, error)
}

type PostgresConfig interface {
	GetUser() string
	GetPassword() string
	GetHost() string
	GetName() string
	GetDialect() string
	GetPort() int
	GetPoolMin() int
	GetPoolMax() int
}
