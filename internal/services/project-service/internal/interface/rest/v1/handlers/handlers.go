package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"project-management-system/internal/project-service/internal/config"
	"project-management-system/internal/project-service/internal/domain/repository/postgres"
	"project-management-system/internal/project-service/internal/service"
	"time"
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

func New(storage Storage, serviceTimeout time.Duration, cfg *config.Config) (*Handlers, error) {
	connection, err := storage.GetConnection()
	if err != nil {
		return nil, err
	}

	//repos
	projectRepo := postgres.NewProjectsRepository(connection)

	//services
	projectService := service.NewProjectService(projectRepo)

	return &Handlers{
		NewProjectsHandler(projectService),
	}, err
}
