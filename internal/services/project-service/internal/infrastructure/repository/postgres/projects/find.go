package projectsPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

func (p *ProjectRepository) findOne(
	ctx context.Context,
	query string,
	args ...interface{},
) (*project.Project, error) {
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = stmt.Close() }()

	found := project.Project{}
	err = stmt.QueryRowContext(ctx, args...).Scan(
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

	if err != nil {
		return nil, ErrProjectNotFound
	}

	return &found, nil
}

func (p *ProjectRepository) FindById(ctx context.Context, id string) (*project.Project, error) {
	query := `SELECT 
    	p.id, p.name, p.description, p.status, p.production_start_at, p.production_end_at, p.created_at, p.updated_at, p.archived_at
		FROM projects AS p WHERE p.id=$1`

	return p.findOne(ctx, query, id)
}
