package projectUserRepoPostgres

import (
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

type ProjectUserRepository struct {
	db           postgres.DB
	txManager    postgres.TxManager
	projectsRepo persistent.ProjectRepository
}

func New(
	db postgres.DB,
	txManager postgres.TxManager,
	projectsRepo persistent.ProjectRepository,
) *ProjectUserRepository {
	return &ProjectUserRepository{
		db:           db,
		txManager:    txManager,
		projectsRepo: projectsRepo,
	}
}

func (pu *ProjectUserRepository) WithTx(tx *sql.Tx) *ProjectUserRepository {
	return New(tx, pu.txManager, pu.projectsRepo)
}
