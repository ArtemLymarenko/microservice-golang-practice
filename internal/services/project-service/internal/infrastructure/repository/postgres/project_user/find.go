package projectUserRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

func (pu *ProjectUserRepository) FindAllProjectMembers(
	ctx context.Context,
	projectId project.Id,
) ([]user.Id, error) {
	query := `SELECT p.user_id FROM projects_users as p WHERE project_id=$1`

	userScan := func(scan postgres.Scanner) (user.Id, error) {
		var found user.Id
		err := scan.Scan(
			&found,
		)

		return found, err
	}

	users, err := postgres.FindMany[user.Id](ctx, pu.db, userScan, query, projectId)
	if err != nil {
		return nil, ErrMembersNotFound
	}

	return users, nil
}

func (pu *ProjectUserRepository) FindAllProjectMembersWithRoles(
	ctx context.Context,
	projectId project.Id,
) ([]valueobject.UserRole, error) {
	query := `SELECT p.role, p.user_id FROM projects_users as p WHERE project_id=$1`

	userRolesScan := func(scan postgres.Scanner) (valueobject.UserRole, error) {
		var found valueobject.UserRole
		err := scan.Scan(
			&found.Role,
			&found.UserId,
		)

		return found, err
	}

	users, err := postgres.FindMany[valueobject.UserRole](ctx, pu.db, userRolesScan, query, projectId)
	if err != nil {
		return nil, ErrMembersWithRoleNotFound
	}

	return users, nil
}
