package handlers

import (
	"github.com/gin-gonic/gin"
)

type ProjectHandlerImpl interface {
	GetProjectById(c *gin.Context)
}

type Handlers struct {
	ProjectHandler ProjectHandlerImpl
}

func New(projectService ProjectsService) *Handlers {
	return &Handlers{
		NewProjectsHandler(projectService),
	}
}
