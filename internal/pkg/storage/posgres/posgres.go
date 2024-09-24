package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log/slog"
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

type Storage struct {
	Db *sql.DB
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

func New(
	ps PostgresConfig,
	env commonconfig.Env,
	log *slog.Logger,
) (*Storage, error) {
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
		return nil, fmt.Errorf("%s: %w", err)
	}

	db.SetMaxIdleConns(ps.GetPoolMin())
	db.SetMaxOpenConns(ps.GetPoolMax())

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("%s: %w", err)
	}

	log.Info("connected to database")

	m, err := migrate.New(
		"file://migrations",
		connectionPath,
	)

	if err != nil {
		_ = db.Close()
		return nil, err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		_ = db.Close()
		return nil, err
	}

	log.Info("migrations created")

	return &Storage{db}, nil
}

func (r *Storage) CloseConnection() error {
	if err := r.Db.Close(); err != nil {
		return err
	}

	return nil
}
