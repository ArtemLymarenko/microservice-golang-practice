package projectsRepoPostgres

import (
	"context"
	sqlStorage "project-management-system/internal/pkg/sql_storage"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
)

func (p *ProjectRepository) scanProject(row sqlStorage.RowScanner) (project.Project, error) {
	var found project.Project
	err := row.Scan(
		&found.Id,
		&found.Name,
		&found.Description,
		&found.Status,
		&found.ProductionStartAt,
		&found.ProductionEndAt,
		&found.CreatedAt,
		&found.UpdatedAt,
		&found.ArchivedAt,
	)

	return found, err
}

func (p *ProjectRepository) FindById(ctx context.Context, id project.Id) (result project.Project, err error) {
	query := `SELECT * FROM projects AS p WHERE p.id=$1 LIMIT 1`

	result, err = sqlStorage.FindOne(ctx, p.db, p.scanProject, query, id)
	if err != nil {
		return result, ErrProjectNotFound
	}

	return result, nil
}

func (p *ProjectRepository) FindByNameMany(
	ctx context.Context,
	name project.Name,
) ([]project.Project, error) {
	query := `SELECT * FROM projects AS p WHERE p.name=$1`

	result, err := sqlStorage.FindMany(ctx, p.db, p.scanProject, query, name)
	if err != nil {
		return nil, ErrProjectsNotFound
	}

	return result, nil
}

func (p *ProjectRepository) FindUserProjects(ctx context.Context, userId user.Id) ([]project.Project, error) {
	query := `
		SELECT 
			p.id, p.name, p.description, p.status, 
			p.production_start_at, p.production_end_at,
			p.created_at, p.updated_at, p.archived_at 
		FROM projects AS p  
		LEFT JOIN projects_users AS pu
		ON p.id = pu.project_id
		WHERE pu.user_id=$1`

	result, err := sqlStorage.FindMany(ctx, p.db, p.scanProject, query, userId)
	if err != nil {
		return nil, ErrProjectsNotFound
	}

	return result, nil
}
