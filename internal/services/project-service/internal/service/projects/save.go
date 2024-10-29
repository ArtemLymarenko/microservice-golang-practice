package projectService

import (
	"context"
	"database/sql"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/valueobject"
	"project-management-system/internal/project-service/internal/service"
)

func (p *ProjectService) CreateWithOwner(
	ctx context.Context,
	ownerId user.Id,
	project project.Project,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	saveProjectWithOwnerTx := func(ctx context.Context, tx *sql.Tx) error {
		projectRepoTx := p.projectsRepo.WithTx(tx)
		projectUserTx := p.projectUserRepo.WithTx(tx)

		if err := projectRepoTx.Save(ctx, project); err != nil {
			return err
		}

		ownerRole := valueobject.UserRole{
			UserId: ownerId,
			Role:   role.Owner,
		}

		return projectUserTx.SaveMemberToProject(ctx, project.Id, ownerRole)
	}

	return p.txManager.Run(ctxWithTimeout, saveProjectWithOwnerTx)
}
