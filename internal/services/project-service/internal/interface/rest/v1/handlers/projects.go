package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-management-system/internal/project-service/internal/domain/model/project"
	"project-management-system/internal/project-service/internal/domain/repository/postgres"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
)

type ProjectsService interface {
	FindById(ctx context.Context, id string) (*project.Project, error)
}

type ProjectsHandlerImpl struct {
	projectService ProjectsService
}

func NewProjectsHandler(projectService ProjectsService) *ProjectsHandlerImpl {
	return &ProjectsHandlerImpl{
		projectService,
	}
}

func (handler *ProjectsHandlerImpl) GetProjectById(c *gin.Context) {
	const param = "id"

	id := c.Param(param)

	found, err := handler.projectService.FindById(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, postgres.ErrProjectNotFound) {
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
