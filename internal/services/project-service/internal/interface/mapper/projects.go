package mapper

import (
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
)

func FromCreateProjectDtoToProjectEntity(request dto.CreateProjectRequest) project.Project {
	return project.Project{
		Name:              project.Name(request.Name),
		Description:       project.Description(request.Description),
		Status:            project.Status(request.Status),
		ProductionStartAt: request.ProductionStartAt,
		ProductionEndAt:   request.ProductionEndAt,
	}
}

func FromAddMemberDtoToUserRole(request dto.AddMemberToProjectRequest) valueobject.UserRole {
	return valueobject.UserRole{
		UserId: user.Id(request.UserId),
		Role:   role.Role(request.Role),
	}
}
