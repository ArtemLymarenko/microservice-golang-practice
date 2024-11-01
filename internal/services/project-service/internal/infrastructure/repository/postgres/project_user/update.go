package projectUserRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
)

func (pu *ProjectUserRepository) ChangeUserRoleInProject(
	ctx context.Context,
	newRole role.Role,
	userId user.Id,
	projectId project.Id,
) error {
	query := `UPDATE projects_users SET role=$1 WHERE user_id=$2 AND project_id=$3`

	_, err := pu.db.ExecContext(ctx, query, newRole, userId, projectId)
	if err != nil {
		return ErrRoleNotChanged
	}

	return nil
}
