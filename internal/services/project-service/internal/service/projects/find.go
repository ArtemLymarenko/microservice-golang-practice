package projectsService

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/service"
)

func (p *ProjectService) FindProjectById(ctx context.Context, id project.Id) (*project.Project, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	return p.projectsRepo.FindById(ctxWithTimeout, id)
}

func (p *ProjectService) FindProjectByName(ctx context.Context, name project.Name) ([]project.Project, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	return p.projectsRepo.FindByNameMany(ctxWithTimeout, name)
}
