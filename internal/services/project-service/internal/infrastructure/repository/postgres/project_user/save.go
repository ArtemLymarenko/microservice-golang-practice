package projectUserRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

func (pu *ProjectUserRepository) SaveMemberToProject(
	ctx context.Context,
	projectId project.Id,
	member valueobject.UserRole,
) error {
	saveProjectUserQuery := `INSERT INTO 
    	projects_users("project_id", "user_id", "role")
		VALUES ($1, $2, $3)`
	_, err := pu.db.ExecContext(
		ctx,
		saveProjectUserQuery,
		postgres.ToNullable(projectId),
		postgres.ToNullable(member.UserId),
		postgres.ToNullable(member.Role),
	)
	if err != nil {
		return ErrSaveMember
	}

	return nil
}
