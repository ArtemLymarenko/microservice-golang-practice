package sqlmapper

import (
	sqlStorage "project-management-system/internal/pkg/sql-storage"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/infrastructure/repository/sql/dto"
)

func FromProjectEntityToRow(from project.Project) (to sqlrow.Project) {
	return sqlrow.Project{
		Id:                string(from.Id),
		Name:              string(from.Name),
		Description:       sqlStorage.ComparableToSqlNull(string(from.Description)),
		Status:            string(from.Status),
		ProductionStartAt: sqlStorage.ComparableToSqlNull(from.ProductionStartAt),
		ProductionEndAt:   sqlStorage.ComparableToSqlNull(from.ProductionEndAt),
		CreatedAt:         sqlStorage.ComparableToSqlNull(from.CreatedAt),
		UpdatedAt:         sqlStorage.ComparableToSqlNull(from.UpdatedAt),
		ArchivedAt:        sqlStorage.ComparableToSqlNull(from.ArchivedAt),
	}
}

func FromUserRoleValueObjToRow(from valueobject.UserRole) (to sqlrow.UserRole) {
	return sqlrow.UserRole{
		UserId: string(from.UserId),
		Role:   string(from.Role),
	}
}
