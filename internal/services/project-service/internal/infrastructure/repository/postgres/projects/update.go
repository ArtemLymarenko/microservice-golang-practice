package projectsRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

func (p *ProjectRepository) UpdateById(ctx context.Context, project project.Project) error {
	updateProjectQuery := `
		UPDATE projects 
		SET 
			name = $1, 
			description = $2, 
			status = $3, 
			production_start_at = $4, 
			production_end_at = $5, 
			created_at = $6,
			updated_at = $7,
			archived_at = $8
		WHERE id = $9;
	`

	_, err := p.db.ExecContext(
		ctx,
		updateProjectQuery,
		project.Name,
		project.Description,
		project.Status,
		project.ProductionStartAt,
		project.ProductionEndAt,
		project.CreatedAt,
		project.UpdatedAt,
		project.ArchivedAt,
		project.Id,
	)

	if err != nil {
		return ErrUpdateProject
	}

	return nil
}
