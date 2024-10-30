package projectService

import (
	"github.com/go-playground/validator/v10"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

type ProjectService struct {
	projectsRepo    persistent.ProjectRepository
	projectUserRepo persistent.ProjectUserRepository
	sqlTxManager    postgres.TxManager
	validator       *validator.Validate
}

func New(
	projectsRepo persistent.ProjectRepository,
	projectUserRepository persistent.ProjectUserRepository,
	sqlTxManager postgres.TxManager,
	validator *validator.Validate,
) *ProjectService {
	return &ProjectService{
		projectsRepo:    projectsRepo,
		projectUserRepo: projectUserRepository,
		sqlTxManager:    sqlTxManager,
		validator:       validator,
	}
}
