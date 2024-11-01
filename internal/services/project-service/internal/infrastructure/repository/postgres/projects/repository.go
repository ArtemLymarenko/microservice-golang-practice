package projectsRepoPostgres

import (
	"database/sql"
	sqlStorage "project-management-system/internal/pkg/sql_storage"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
)

type ProjectRepository struct {
	db sqlStorage.DB
}

func New(db sqlStorage.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (p *ProjectRepository) WithTx(tx *sql.Tx) persistent.ProjectRepository {
	return New(tx)
}
