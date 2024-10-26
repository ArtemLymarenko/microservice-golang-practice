package projectsPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

func (p *ProjectRepository) Save(ctx context.Context, project project.Project) error {
	saveProjectQuery := `INSERT INTO 
    	projects(id, name, description, status, production_start_at, production_end_at, created_at, updated_at, archived_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := p.db.ExecContext(
		ctx,
		saveProjectQuery,
		project.Id,
		project.Name,
		project.Description,
		project.Status,
		project.ProductionStartAt,
		project.ProductionEndAt,
		project.CreatedAt,
		project.UpdatedAt,
		project.ArchivedAt,
	)

	if err != nil {
		return ErrSaveProject
	}

	return err
}
