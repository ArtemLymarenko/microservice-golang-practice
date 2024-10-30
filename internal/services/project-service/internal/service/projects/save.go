package projectService

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

	projectId := uuid.New().String()
	proj.SetId(project.Id(projectId))

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

		return projectUserTx.SaveMemberToProject(ctx, proj.Id, ownerRole)
	}

	return p.sqlTxManager.Run(ctxWithTimeout, saveProjectWithOwnerTx)
}
