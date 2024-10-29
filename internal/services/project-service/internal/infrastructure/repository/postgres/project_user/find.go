package projectUserRepoPostgres

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
)

func (pu *ProjectUserRepository) FindAllProjectMembers(
	ctx context.Context,
	projectId project.Id,
) ([]user.Id, error) {
	query := `SELECT p.user_id FROM projects_users as p WHERE project_id=$1`

	userScan := func(row postgres.RowScanner) (user.Id, error) {
		var found user.Id
		err := row.Scan(
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

	userRolesScan := func(row postgres.RowScanner) (valueobject.UserRole, error) {
		var found valueobject.UserRole
		err := row.Scan(
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

func (pu *ProjectUserRepository) FindUserRoleByProject(
	ctx context.Context,
	userId user.Id,
	projectId project.Id,
) (result role.Role, err error) {
	query := `SELECT p.role FROM projects_users as p WHERE project_id=$1 AND user_id=$2 LIMIT 1`

	roleScan := func(row postgres.RowScanner) (found role.Role, err error) {
		err = row.Scan(
			&found,
		)

		return found, err
	}

	result, err = postgres.FindOne[role.Role](ctx, pu.db, roleScan, query, userId, projectId)
	if err != nil {
		return result, ErrMembersWithRoleNotFound
	}

	return result, nil
}
