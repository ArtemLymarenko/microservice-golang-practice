package persistent

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

type ProjectRepository interface {
	FindById(ctx context.Context, id project.Id) (*project.Project, error)
	FindByNameMany(ctx context.Context, name project.Name) ([]project.Project, error)
	Save(ctx context.Context, proj project.Project) error
	UpdateById(ctx context.Context, project project.Project) error
	DeleteById(ctx context.Context, projectId project.Id) error
	WithTx(tx *sql.Tx) ProjectRepository
}
