package mapper

import (
	"project-management-system/internal/project-service/internal/domain/model/project"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
)

func FromProjectRequestDtoToProject(request dto.CreateProjectRequest) project.Project {
	return project.Project{
		Name:        request.Name,
		Description: request.Description,
	}
}
