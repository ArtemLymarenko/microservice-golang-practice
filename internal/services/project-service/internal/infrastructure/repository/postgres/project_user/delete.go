package projectUserRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

func (pu *ProjectUserRepository) DeleteUserFromProject(ctx context.Context, projectId project.Id) error {
	query := `DELETE FROM  projects_users WHERE project_id=$1`
	_, err := pu.db.ExecContext(ctx, query, projectId)
	if err != nil {
		return ErrDeleteMember
	}

	return nil
}
