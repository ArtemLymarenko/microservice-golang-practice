package dto

import "project-management-system/internal/project-service/internal/domain/entity/project"

type CreateProjectResponse struct {
	Id string `json:"id,omitempty"`
}

type GetProjectByIdResponse struct {
	Project project.Project `json:"project"`
}
