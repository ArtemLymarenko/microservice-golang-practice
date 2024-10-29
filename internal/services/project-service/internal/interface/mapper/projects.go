package mapper

import (
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
)

func FromProjectRequestDtoToProject(request dto.CreateProjectRequest) project.Project {
	return project.Project{
		Name:              project.Name(request.Name),
		Description:       project.Description(request.Description),
		ProductionStartAt: request.ProductionStartAt,
		ProductionEndAt:   request.ProductionEndAt,
	}
}
