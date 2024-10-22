package postgres

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/model/project"
)

type ProjectsRepository struct {
	db *sql.DB
}

func NewProjectsRepository(db *sql.DB) *ProjectsRepository {
	return &ProjectsRepository{db}
}

func (p *ProjectsRepository) findOne(
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
		return nil, ErrQueryRowWithContext
	}

	return &found, nil
}

func (p *ProjectsRepository) FindById(ctx context.Context, id string) (*project.Project, error) {
	query := `SELECT 
    	p.id, p.name, p.description, p.status, p.production_start_at, p.production_end_at, p.created_at, p.updated_at, p.archived_at
		FROM projects AS p WHERE p.id=$1`

	found, err := p.findOne(ctx, query, id)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	return found, nil
}

func (p *ProjectsRepository) SaveByUser(ctx context.Context, userId string, project project.Project) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	saveProjectQuery := `INSERT INTO 
    	projects(id, name, description, status, production_start_at, production_end_at, created_at, updated_at, archived_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = tx.ExecContext(
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
		return ErrCommitTrx
	}

	saveProjectUserQuery := `INSERT INTO 
    	projects_users(project_id, user_id)
		VALUES ($1, $2)`

	_, err = tx.ExecContext(
		ctx,
		saveProjectUserQuery,
		project.Id,
		userId,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return ErrCommitTrx
	}

	return nil
}
