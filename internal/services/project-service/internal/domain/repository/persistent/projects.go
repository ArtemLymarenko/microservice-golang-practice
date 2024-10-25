package persistent

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

type ProjectRepository interface {
	FindById(ctx context.Context, id string) (*project.Project, error)
	Save(ctx context.Context, project project.Project) error
}
