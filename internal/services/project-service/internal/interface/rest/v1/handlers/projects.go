package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	projectsPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/projects"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
)

type ProjectsService interface {
	FindProjectById(ctx context.Context, id project.Id) (project.Project, error)
}

type ProjectsHandlerImpl struct {
	projectsService ProjectsService
}

func NewProjectsHandler(projectService ProjectsService) *ProjectsHandlerImpl {
	return &ProjectsHandlerImpl{
		projectService,
	}
}

func (handler *ProjectsHandlerImpl) GetProjectById(c *gin.Context) {
	const param = "id"

	id := project.Id(c.Param(param))
	ctx := c.Request.Context()

	found, err := handler.projectsService.FindProjectById(ctx, id)
	if err != nil {
		if errors.Is(err, projectsPostgres.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, dto.NewResponseErr(err))
			return
		}

		c.JSON(http.StatusInternalServerError, dto.NewResponseErr(err))
		return
	}

	c.JSON(http.StatusOK, dto.GetProjectByIdResponse{
		Project: found,
	})
}
