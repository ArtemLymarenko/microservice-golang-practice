package persistent

import (
	"context"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
)

type ProjectUserRepository interface {
	SaveProjectWithOwner(
		ctx context.Context,
		ownerId user.Id,
		project project.Project,
	) error
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
	DeleteProjectMember(ctx context.Context, projectId project.Id) error
}
