package projectsRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

func (p *ProjectRepository) DeleteById(ctx context.Context, projectId project.Id) error {
	deleteProjectQuery := `DELETE FROM projects WHERE id=$1`

	_, err := p.db.ExecContext(
		ctx,
		deleteProjectQuery,
		projectId,
	)

	if err != nil {
		return ErrDeleteProject
	}

	return nil
}
