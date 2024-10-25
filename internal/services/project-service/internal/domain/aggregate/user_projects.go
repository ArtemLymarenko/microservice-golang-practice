package aggregate

import (
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
)

type UserProjects struct {
	UserId   user.Id
	Projects []project.Project
}
