package projectUserRepoPostgres

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
	projectsRepoPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/projects"
)

type ProjectsRepo interface {
	Save(ctx context.Context, project project.Project) error
	WithTx(tx *sql.Tx) *projectsRepoPostgres.ProjectRepository
}

type ProjectUserRepository struct {
	db           postgres.DB
	txManager    postgres.TxManager
	projectsRepo ProjectsRepo
}

func New(
	db postgres.DB,
	txManager postgres.TxManager,
	projectsRepo ProjectsRepo,
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
