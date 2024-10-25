package postgresTx

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/repository/postgres"
)

type ProjectsRepo interface {
	Save(ctx context.Context, project project.Project) error
	WithTx(tx *sql.Tx) *postgres.ProjectRepository
}

type ProjectUserRepo interface {
	Save(ctx context.Context, projectId string, userId string) error
	WithTx(tx *sql.Tx) *postgres.ProjectUserRepository
}

type ProjectUserRepository struct {
	db              *sql.DB
	projectsRepo    ProjectsRepo
	projectUserRepo ProjectUserRepo
}

func New(
	db *sql.DB,
	projectsRepo ProjectsRepo,
	projectUserRepo ProjectUserRepo,
) *ProjectUserRepository {
	return &ProjectUserRepository{
		db:              db,
		projectsRepo:    projectsRepo,
		projectUserRepo: projectUserRepo,
	}
}

func (pu *ProjectUserRepository) SaveProjectWithUser(
	ctx context.Context,
	userId string,
	project project.Project,
) error {
	tx, err := pu.db.BeginTx(ctx, nil)
	if err != nil {
		return ErrCommitTrx
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = pu.projectsRepo.WithTx(tx).Save(ctx, project)
	if err != nil {
		return err
	}

	err = pu.projectUserRepo.WithTx(tx).Save(ctx, project.Id, userId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return ErrCommitTrx
	}

	return nil
}
