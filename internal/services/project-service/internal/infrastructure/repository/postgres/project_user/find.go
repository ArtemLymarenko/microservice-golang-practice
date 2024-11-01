package projectUserRepoPostgres

import (
	"context"
	"project-management-system/internal/pkg/sqlStorage"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
)

func (pu *ProjectUserRepository) FindAllProjectMembers(
	ctx context.Context,
	projectId project.Id,
) ([]user.Id, error) {
	query := `SELECT p.user_id FROM projects_users as p WHERE project_id=$1`

	scanUser := func(row sqlStorage.RowScanner) (userId user.Id, err error) {
		err = row.Scan(
			&userId,
		)

		return userId, err
	}

	userIds, err := sqlStorage.FindMany(ctx, pu.db, scanUser, query, projectId)
	if err != nil {
		return nil, ErrMembersNotFound
	}

	return userIds, nil
}

func (pu *ProjectUserRepository) FindAllProjectMembersWithRoles(
	ctx context.Context,
	projectId project.Id,
) ([]valueobject.UserRole, error) {
	query := `SELECT p.role, p.user_id FROM projects_users as p WHERE project_id=$1`

	scanUserRoles := func(row sqlStorage.RowScanner) (result valueobject.UserRole, err error) {
		err = row.Scan(
			&result.Role,
			&result.UserId,
		)

		return result, err
	}

	usersWithRoles, err := sqlStorage.FindMany(ctx, pu.db, scanUserRoles, query, projectId)
	if err != nil {
		return nil, ErrMembersWithRoleNotFound
	}

	return usersWithRoles, nil
}

func (pu *ProjectUserRepository) FindUserRoleByProject(
	ctx context.Context,
	userId user.Id,
	projectId project.Id,
) (foundRole role.Role, err error) {
	query := `SELECT p.role FROM projects_users as p WHERE project_id=$1 AND user_id=$2 LIMIT 1`

	scanRole := func(row sqlStorage.RowScanner) (result role.Role, err error) {
		err = row.Scan(
			&result,
		)

		return result, err
	}

	foundRole, err = sqlStorage.FindOne(ctx, pu.db, scanRole, query, userId, projectId)
	if err != nil {
		return foundRole, ErrMemberWithRoleNotFound
	}

	return foundRole, nil
}
