package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
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

func New(projectService ProjectsService) *Handlers {
	return &Handlers{
		NewProjectsHandler(projectService),
	}
}
