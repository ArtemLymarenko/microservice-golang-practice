package projectsService

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/service"
)

func (p *ProjectService) UpdateProject(
	ctx context.Context,
	proj project.Project,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	if err := proj.Validate(p.validator); err != nil {
		return err
	}

	return p.projectsRepo.Update(ctxWithTimeout, proj)
}
