package persistent

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
)

type ProjectRepository interface {
	FindById(ctx context.Context, id project.Id) (project.Project, error)
	FindByNameMany(ctx context.Context, name project.Name) ([]project.Project, error)
	FindUserProjects(ctx context.Context, userId user.Id) ([]project.Project, error)
	Save(ctx context.Context, proj project.Project) error
	Update(ctx context.Context, project project.Project) error
	DeleteById(ctx context.Context, projectId project.Id) error
	EnrichProjects(ctx context.Context, projectIds []project.Id) (projects []project.Project, err error)
	WithTx(tx *sql.Tx) ProjectRepository
}
