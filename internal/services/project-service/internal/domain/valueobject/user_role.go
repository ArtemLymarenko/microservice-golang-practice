package valueobject

import (
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
)

type UserRole struct {
	UserId user.Id
	Role   role.Role
}
