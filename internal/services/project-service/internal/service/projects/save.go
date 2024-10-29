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
	"time"
)

func (p *ProjectService) CreateWithOwner(
	ctx context.Context,
	ownerId user.Id,
	projectToCreate project.Project,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	projectId := project.Id(uuid.New().String())
	projectToCreate.SetId(projectId)
	projectToCreate.SetStatus(project.StatusIdle)
	projectToCreate.SetCreatedAt(time.Now())
	projectToCreate.SetUpdatedAt(time.Now())

	if err := projectToCreate.Validate(p.validator); err != nil {
		return err
	}

	saveProjectWithOwnerTx := func(ctx context.Context, tx *sql.Tx) error {
		projectRepoTx := p.projectsRepo.WithTx(tx)
		projectUserTx := p.projectUserRepo.WithTx(tx)

		if err := projectRepoTx.Save(ctx, projectToCreate); err != nil {
			return err
		}

		ownerRole := valueobject.UserRole{
			UserId: ownerId,
			Role:   role.Owner,
		}

		return projectUserTx.SaveMemberToProject(ctx, projectToCreate.Id, ownerRole)
	}

	return p.txManager.Run(ctxWithTimeout, saveProjectWithOwnerTx)
}
