package service

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/model/project"
)

type ProjectsRepository interface {
	FindById(ctx context.Context, id string) (*project.Project, error)
	SaveByUser(ctx context.Context, userId string, project project.Project) error
}

type ProjectService struct {
	projectsRepo ProjectsRepository
}

func NewProjectService(projectsRepo ProjectsRepository) *ProjectService {
	return &ProjectService{projectsRepo}
}

func (p *ProjectService) FindById(ctx context.Context, id string) (*project.Project, error) {
	return p.projectsRepo.FindById(ctx, id)
}

func (p *ProjectService) SaveByUser(ctx context.Context, userId string, project project.Project) error {
	return p.projectsRepo.SaveByUser(ctx, userId, project)
}
