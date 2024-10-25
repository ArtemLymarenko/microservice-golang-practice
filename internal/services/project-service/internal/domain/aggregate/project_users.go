package aggregate

import (
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/valueobject"
)

type ProjectUsers struct {
	Project project.Project
	UserIds []valueobject.UserRole
}
