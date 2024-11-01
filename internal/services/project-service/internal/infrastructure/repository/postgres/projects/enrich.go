package projectsRepoPostgres

import (
	"context"
	sqlStorage "project-management-system/internal/pkg/sql_storage"
	"project-management-system/internal/project-service/internal/domain/entity/project"
)

func (p *ProjectRepository) EnrichProjects(
	ctx context.Context,
	projectIds []project.Id,
) (projects []project.Project, err error) {
	query := `SELECT * FROM projects AS p WHERE p.id=ANY($1)`

	projects, err = sqlStorage.FindMany(ctx, p.db, p.scanProject, query, projectIds)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
