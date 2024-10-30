package projectsRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	postgresMapper "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/mapper"
)

func (p *ProjectRepository) Save(ctx context.Context, proj project.Project) error {
	projectToStore := postgresMapper.FromProjectEntityToPostgres(proj)

	saveProjectQuery := `INSERT INTO 
    	projects(id, name, description, status, production_start_at, production_end_at, created_at, updated_at, archived_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := p.db.ExecContext(
		ctx,
		saveProjectQuery,
		projectToStore.Id,
		projectToStore.Name,
		projectToStore.Description,
		projectToStore.Status,
		projectToStore.ProductionStartAt,
		projectToStore.ProductionEndAt,
		projectToStore.CreatedAt,
		projectToStore.UpdatedAt,
		projectToStore.ArchivedAt,
	)

	if err != nil {
		return ErrSaveProject
	}

	return err
}
