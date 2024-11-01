package projectUserRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	sqlmapper "project-management-system/internal/project-service/internal/infrastructure/repository/sql/mapper"
)

func (pu *ProjectUserRepository) AddUserToProject(
	ctx context.Context,
	projectId project.Id,
	member valueobject.UserRole,
) error {
	memberToStore := sqlmapper.FromUserRoleValueObjToRow(member)

	saveProjectUserQuery := `INSERT INTO 
    	projects_users("project_id", "user_id", "role")
		VALUES ($1, $2, $3)`

	_, err := pu.db.ExecContext(
		ctx,
		saveProjectUserQuery,
		string(projectId),
		memberToStore.UserId,
		memberToStore.Role,
	)
	if err != nil {
		return ErrSaveMember
	}

	return nil
}
