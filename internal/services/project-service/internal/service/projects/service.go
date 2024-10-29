package projectService

import (
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

type ProjectService struct {
	projectsRepo    persistent.ProjectRepository
	projectUserRepo persistent.ProjectUserRepository
	txManager       postgres.TxManager
}

func New(
	projectsRepo persistent.ProjectRepository,
	projectUserRepository persistent.ProjectUserRepository,
	txManager postgres.TxManager,
) *ProjectService {
	return &ProjectService{
		projectsRepo:    projectsRepo,
		projectUserRepo: projectUserRepository,
		txManager:       txManager,
	}
}
