package projectsRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

func (p *ProjectRepository) scanProject(row postgres.RowScanner) (project.Project, error) {
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

func (p *ProjectRepository) FindById(ctx context.Context, id project.Id) (*project.Project, error) {
	query := `SELECT * FROM projects AS p WHERE p.id=$1 LIMIT 1`

	found, err := postgres.FindOne[project.Project](ctx, p.db, p.scanProject, query, id)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	return &found, nil
}

func (p *ProjectRepository) FindByNameMany(
	ctx context.Context,
	name project.Name,
) ([]project.Project, error) {
	query := `SELECT * FROM projects AS p WHERE p.name=$1`

	found, err := postgres.FindMany[project.Project](ctx, p.db, p.scanProject, query, name)
	if err != nil {
		return nil, ErrProjectsNotFound
	}

	return found, nil
}
