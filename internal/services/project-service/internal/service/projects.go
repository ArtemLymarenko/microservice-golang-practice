package service

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

type ProjectRepository interface {
	FindById(ctx context.Context, id string) (*project.Project, error)
}

type ProjectUserTxRepository interface {
	SaveProjectWithUser(
		ctx context.Context,
		userId string,
		project project.Project,
	) error
}

type ProjectService struct {
	projectsRepo            ProjectRepository
	projectUserTxRepository ProjectUserTxRepository
}

func NewProjectService(
	projectsRepo ProjectRepository,
	projectUserTxRepository ProjectUserTxRepository,
) *ProjectService {
	return &ProjectService{
		projectsRepo:            projectsRepo,
		projectUserTxRepository: projectUserTxRepository,
	}
}

func (p *ProjectService) GetById(ctx context.Context, id string) (*project.Project, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	return p.projectsRepo.FindById(ctxWithTimeout, id)
}

func (p *ProjectService) AddProjectWithCreator(
	ctx context.Context,
	creatorId string,
	project project.Project,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	//Some validation

	err := p.projectUserTxRepository.SaveProjectWithUser(ctxWithTimeout, creatorId, project)
	if err != nil {
		return err
	}

	return nil
}
