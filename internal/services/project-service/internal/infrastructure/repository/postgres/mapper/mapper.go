package postgresMapper

import (
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres/dto"
)

func FromProjectEntityToPostgres(from project.Project) (to postgresdto.Project) {
	return postgresdto.Project{
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
