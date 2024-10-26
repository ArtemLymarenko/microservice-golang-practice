package projectUserPostgres

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
	projectsPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/projects"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres/tx"
)

type ProjectsRepo interface {
	Save(ctx context.Context, project project.Project) error
	WithTx(tx *sql.Tx) *projectsPostgres.ProjectRepository
}

type ProjectUserRepository struct {
	db           postgres.DB
	txManager    tx.Manager
	projectsRepo ProjectsRepo
}

func New(
	db postgres.DB,
	txManager tx.Manager,
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
