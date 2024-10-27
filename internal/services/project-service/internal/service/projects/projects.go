package projectService

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
	"project-management-system/internal/project-service/internal/service"
)

type ProjectService struct {
	projectsRepo            persistent.ProjectRepository
	projectUserTxRepository persistent.ProjectUserRepository
}

func New(
	projectsRepo persistent.ProjectRepository,
	projectUserRepository persistent.ProjectUserRepository,
) *ProjectService {
	return &ProjectService{
		projectsRepo:            projectsRepo,
		projectUserTxRepository: projectUserRepository,
	}
}

func (p *ProjectService) GetById(ctx context.Context, id project.Id) (*project.Project, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	return p.projectsRepo.FindById(ctxWithTimeout, id)
}

func (p *ProjectService) AddProjectWithOwner(
	ctx context.Context,
	ownerId user.Id,
	project project.Project,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	//Some validation

	err := p.projectUserTxRepository.SaveProjectWithOwner(ctxWithTimeout, ownerId, project)
	if err != nil {
		return err
	}

	return nil
}
