package projectsPostgres

import (
	"context"
	"database/sql"
	"errors"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

var (
	ErrProjectNotFound = errors.New("project was not found")
	ErrSaveProject     = errors.New("failed to save project")
	ErrDeleteProject   = errors.New("failed to delete project")
	ErrUpdateProject   = errors.New("failed to update project")
)

type ProjectRepository struct {
	db postgres.DB
}

func New(db postgres.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (p *ProjectRepository) WithTx(tx *sql.Tx) *ProjectRepository {
	return New(tx)
}

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
