package projectsRepoPostgres

import (
	"database/sql"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

type ProjectRepository struct {
	db postgres.DB
}

func New(db postgres.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (p *ProjectRepository) WithTx(tx *sql.Tx) *ProjectRepository {
	return New(tx)
}
