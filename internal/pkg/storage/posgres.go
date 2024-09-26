package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"project-management-system/internal/pkg/config"
)

const (
	sslDisable    = "disable"
	sslVerifyFull = "verify-full"
)

func getSSLConfig(env commonconfig.Env) string {
	var sslConfig string
	if env == commonconfig.EnvLocal {
		sslConfig = sslDisable
	} else {
		sslConfig = sslVerifyFull
	}

	return sslConfig
}

type postgres struct {
	db *sql.DB
}

func NewPostgres(
	ps PostgresConfig,
	env commonconfig.Env,
) (Postgres, error) {
	connectionPath := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		ps.GetUser(),
		ps.GetPassword(),
		ps.GetHost(),
		ps.GetPort(),
		ps.GetName(),
		getSSLConfig(env),
	)

	db, err := sql.Open(ps.GetDialect(), connectionPath)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(ps.GetPoolMin())
	db.SetMaxOpenConns(ps.GetPoolMax())

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, ErrHealthCheckPostgres
	}

	logrus.Info("connected to database")

	m, err := migrate.New(
		"file://migrations",
		connectionPath,
	)

	if err != nil {
		_ = db.Close()
		return nil, ErrCreateMigrationPostgres
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		_ = db.Close()
		return nil, ErrMigratePostgres
	}

	logrus.Info("migrations created")

	return &postgres{db}, nil
}

func (p *postgres) CloseConnection() error {
	if err := p.db.Close(); err != nil {
		return ErrClosePostgresConnection
	}

	return nil
}

func (p *postgres) GetConnection() (*sql.DB, error) {
	if p.db != nil {
		return p.db, nil
	}

	return nil, ErrGetPostgresConnection
}
