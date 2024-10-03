package postgres

import (
	"context"
	"database/sql"
	"errors"
	"project-management-system/internal/project-service/internal/domain/model"
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
) (*model.Project, error) {
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	project := model.Project{}
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&project.Id,
		&project.Name,
		&project.Description,
		&project.Status,
		&project.UpdatedAt,
		&project.ProductionStartAt,
		&project.ProductionEndAt,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.ArchivedAt,
	)

	if err != nil {
		return nil, errors.Join(ErrQueryRowWithContext, err)
	}

	return &project, nil
}
