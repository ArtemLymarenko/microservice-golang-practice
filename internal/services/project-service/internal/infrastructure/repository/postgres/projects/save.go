package projectsRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

func (p *ProjectRepository) Save(ctx context.Context, projectToSave project.Project) error {
	saveProjectQuery := `INSERT INTO 
    	projects(id, name, description, status, production_start_at, production_end_at, created_at, updated_at, archived_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := p.db.ExecContext(
		ctx,
		saveProjectQuery,
		projectToSave.Id,
		projectToSave.Name,
		projectToSave.Description,
		projectToSave.Status,
		postgres.ToNullable(projectToSave.ProductionStartAt),
		postgres.ToNullable(projectToSave.ProductionEndAt),
		projectToSave.CreatedAt,
		projectToSave.UpdatedAt,
		postgres.ToNullable(projectToSave.ArchivedAt),
	)

	if err != nil {
		return ErrSaveProject
	}

	return err
}
