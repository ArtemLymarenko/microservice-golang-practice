package projectsService

import (
	sqlStorage "project-management-system/internal/pkg/sql_storage"
	"project-management-system/internal/project-service/internal/domain/entity"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
)

type ProjectService struct {
	projectsRepo    persistent.ProjectRepository
	projectUserRepo persistent.ProjectUserRepository
	sqlTxManager    sqlStorage.TxManager
	validator       entity.Validator
}

func New(
	projectsRepo persistent.ProjectRepository,
	projectUserRepository persistent.ProjectUserRepository,
	sqlTxManager sqlStorage.TxManager,
	validator entity.Validator,
) *ProjectService {
	return &ProjectService{
		projectsRepo:    projectsRepo,
		projectUserRepo: projectUserRepository,
		sqlTxManager:    sqlTxManager,
		validator:       validator,
	}
}
