package projectService

import (
	"project-management-system/internal/project-service/internal/domain/entity"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

type ProjectService struct {
	projectsRepo    persistent.ProjectRepository
	projectUserRepo persistent.ProjectUserRepository
	sqlTxManager    postgres.TxManager
	validator       entity.Validator
}

func New(
	projectsRepo persistent.ProjectRepository,
	projectUserRepository persistent.ProjectUserRepository,
	sqlTxManager postgres.TxManager,
	validator entity.Validator,
) *ProjectService {
	return &ProjectService{
		projectsRepo:    projectsRepo,
		projectUserRepo: projectUserRepository,
		sqlTxManager:    sqlTxManager,
		validator:       validator,
	}
}
