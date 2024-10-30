package projectsRepoPostgres

import (
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

type ProjectRepository struct {
	db postgres.DB
}

func New(db postgres.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (p *ProjectRepository) WithTx(tx *sql.Tx) persistent.ProjectRepository {
	return New(tx)
}
