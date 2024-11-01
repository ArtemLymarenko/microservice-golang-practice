package projectsService

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/service"
)

func (p *ProjectService) FindProjectById(ctx context.Context, id project.Id) (project.Project, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	return p.projectsRepo.FindById(ctxWithTimeout, id)
}

func (p *ProjectService) FindProjectByName(ctx context.Context, name project.Name) ([]project.Project, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	return p.projectsRepo.FindByNameMany(ctxWithTimeout, name)
}

func (p *ProjectService) FindUserProjects(ctx context.Context, userId user.Id) ([]project.Project, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	projects, err := p.projectsRepo.FindUserProjects(ctxWithTimeout, userId)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
