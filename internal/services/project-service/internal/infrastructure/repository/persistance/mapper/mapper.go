package sqlmapper

import (
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/infrastructure/repository/persistance/dto"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

func FromProjectEntityToRow(from project.Project) (to sqlrow.Project) {
	return sqlrow.Project{
		Id:                string(from.Id),
		Name:              string(from.Name),
		Description:       postgres.ToNullable(string(from.Description)),
		Status:            string(from.Status),
		ProductionStartAt: postgres.ToNullable(from.ProductionStartAt),
		ProductionEndAt:   postgres.ToNullable(from.ProductionEndAt),
		CreatedAt:         from.CreatedAt,
		UpdatedAt:         from.UpdatedAt,
		ArchivedAt:        postgres.ToNullable(from.ArchivedAt),
	}
}

func FromUserRoleValueObjToRow(from valueobject.UserRole) (to sqlrow.UserRole) {
	return sqlrow.UserRole{
		UserId: string(from.UserId),
		Role:   string(from.Role),
	}
}
