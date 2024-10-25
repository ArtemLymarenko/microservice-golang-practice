package projectUserPostgres

import (
	"context"
	"database/sql"
	"errors"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres/projects"
)

var (
	ErrSaveProjectWithUser = errors.New("failed to save project with user")
)

type ProjectsRepo interface {
	Save(ctx context.Context, project project.Project) error
	WithTx(tx *sql.Tx) *projectsPostgres.ProjectRepository
}

type ProjectUserRepository struct {
	db           *sql.DB
	projectsRepo ProjectsRepo
}

func New(
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
		return postgres.ErrFinishTx
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
		return postgres.ErrFinishTx
	}

	return nil
}

func (pu *ProjectUserRepository) AddUsersToProject(
	ctx context.Context,
	projectId project.Id,
	userIds []user.Id,
) error {
	return nil
}

func (pu *ProjectUserRepository) FindProjectUsers(
	ctx context.Context,
	projectId project.Id,
) []user.Id {
	return nil
}

func (pu *ProjectUserRepository) FindUserRoleInProject(
	ctx context.Context,
	userId user.Id,
	projectId project.Id,
) role.Role {
	return role.Role("")
}
