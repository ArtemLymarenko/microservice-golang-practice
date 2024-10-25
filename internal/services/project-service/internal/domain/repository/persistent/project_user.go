package persistent

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

type ProjectUserRepository interface {
	SaveProjectWithUser(
		ctx context.Context,
		userId string,
		project project.Project,
	) error
}
