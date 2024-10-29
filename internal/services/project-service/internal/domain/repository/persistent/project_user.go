package persistent

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	projectUserRepoPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/project_user"
)

type ProjectUserRepository interface {
	SaveMemberToProject(
		ctx context.Context,
		projectId project.Id,
		member valueobject.UserRole,
	) error
	FindAllProjectMembers(
		ctx context.Context,
		projectId project.Id,
	) ([]user.Id, error)
	FindAllProjectMembersWithRoles(
		ctx context.Context,
		projectId project.Id,
	) ([]valueobject.UserRole, error)
	FindUserRoleByProject(
		ctx context.Context,
		userId user.Id,
		projectId project.Id,
	) (result role.Role, err error)
	DeleteProjectMember(ctx context.Context, projectId project.Id) error
	WithTx(tx *sql.Tx) *projectUserRepoPostgres.ProjectUserRepository
}
