package postgres

import (
	"context"
	"database/sql"
	"errors"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

var (
	ErrSaveProjectWithUser = errors.New("failed to save project with user")
)

type ProjectsRepo interface {
	Save(ctx context.Context, project project.Project) error
	WithTx(tx *sql.Tx) *ProjectRepository
}

type ProjectUserRepository struct {
	db           *sql.DB
	projectsRepo ProjectsRepo
}

func NewProjectUserRepository(
	db *sql.DB,
	projectsRepo ProjectsRepo,
) *ProjectUserRepository {
	return &ProjectUserRepository{
		db:           db,
		projectsRepo: projectsRepo,
	}
}

func (pu *ProjectUserRepository) SaveProjectWithUser(
	ctx context.Context,
	userId string,
	project project.Project,
) error {
	tx, err := pu.db.BeginTx(ctx, nil)
	if err != nil {
		return ErrFinishTx
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

	saveProjectUserQuery := `INSERT INTO 
    	projects_users("project_id", "user_id")
		VALUES ($1, $2)`
	_, err = tx.ExecContext(ctx, saveProjectUserQuery, project.Id, userId)
	if err != nil {
		return ErrSaveProjectWithUser
	}

	err = tx.Commit()
	if err != nil {
		return ErrFinishTx
	}

	return nil
}

func (pu *ProjectUserRepository) FindUserRole(ctx context.Context, projectId project.Id) {

}
