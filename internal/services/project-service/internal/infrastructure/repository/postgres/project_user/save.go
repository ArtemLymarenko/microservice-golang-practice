package projectUserRepoPostgres

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
)

func (pu *ProjectUserRepository) SaveProjectWithOwner(
	ctx context.Context,
	ownerId user.Id,
	project project.Project,
) error {
	transaction := func(ctx context.Context, tx *sql.Tx) error {
		projectRepoTx := pu.projectsRepo.WithTx(tx)
		projectUserTx := pu.WithTx(tx)

		if err := projectRepoTx.Save(ctx, project); err != nil {
			return err
		}

		ownerRole := valueobject.UserRole{
			UserId: ownerId,
			Role:   role.Owner,
		}

		if err := projectUserTx.SaveMemberToProject(ctx, project.Id, ownerRole); err != nil {
			return err
		}

		return nil
	}

	return pu.txManager.Run(ctx, transaction)
}

func (pu *ProjectUserRepository) SaveMemberToProject(
	ctx context.Context,
	projectId project.Id,
	member valueobject.UserRole,
) error {
	saveProjectUserQuery := `INSERT INTO 
    	projects_users("project_id", "user_id", "role")
		VALUES ($1, $2, $3)`
	_, err := pu.db.ExecContext(ctx, saveProjectUserQuery, projectId, member.UserId, member.Role)
	if err != nil {
		return ErrSaveMember
	}

	return nil
}
