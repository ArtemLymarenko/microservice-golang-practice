package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres/project_user"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres/projects"
	"project-management-system/internal/project-service/internal/service/projects"
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
	txManager := postgres.NewTxManager(connection)
	projectRepo := projectsRepoPostgres.New(connection)
	projectUserRepo := projectUserRepoPostgres.New(connection, txManager, projectRepo)

	//services
	projectServ := projectService.New(projectRepo, projectUserRepo)

	return &Handlers{
		NewProjectsHandler(projectServ),
	}, err
}
