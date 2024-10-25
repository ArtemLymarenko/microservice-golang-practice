package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
	"project-management-system/internal/project-service/internal/service"
)

type Storage interface {
	GetConnection() (*sql.DB, error)
}

type ProjectHandlerIml interface {
	GetProjectById(c *gin.Context)
}

type Handlers struct {
	ProjectHandler ProjectHandlerIml
}

func New(storage Storage) (*Handlers, error) {
	connection, err := storage.GetConnection()
	if err != nil {
		return nil, err
	}

	//repos
	projectRepo := postgres.NewProjectRepository(connection)
	projectUserRepo := postgres.NewProjectUserRepository(connection, projectRepo)

	//services
	projectService := service.NewProjectService(projectRepo, projectUserRepo)

	return &Handlers{
		NewProjectsHandler(projectService),
	}, err
}
