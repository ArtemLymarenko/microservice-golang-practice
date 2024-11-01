package projectsService

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/service"
)

func (p *ProjectService) CreateProjectWithOwner(
	ctx context.Context,
	ownerId user.Id,
	proj project.Project,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	projectId := project.Id(uuid.New().String())
	proj.SetId(projectId)

	if err := proj.Validate(p.validator); err != nil {
		return err
	}

	saveProjectWithOwnerTx := func(ctx context.Context, tx *sql.Tx) error {
		projectRepoTx := p.projectsRepo.WithTx(tx)
		projectUserTx := p.projectUserRepo.WithTx(tx)

		if err := projectRepoTx.Save(ctx, proj); err != nil {
			return err
		}

		ownerRole := valueobject.UserRole{
			UserId: ownerId,
			Role:   role.Owner,
		}

		return projectUserTx.AddUserToProject(ctx, proj.Id, ownerRole)
	}

	return p.sqlTxManager.Run(ctxWithTimeout, saveProjectWithOwnerTx)
}

func (p *ProjectService) AddMemberToProject(
	ctx context.Context,
	projectId project.Id,
	memberWithRole valueobject.UserRole,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	if err := role.Validate(memberWithRole.Role); err != nil {
		return err
	}

	if err := p.projectUserRepo.AddUserToProject(ctxWithTimeout, projectId, memberWithRole); err != nil {
		return err
	}

	return nil
}
