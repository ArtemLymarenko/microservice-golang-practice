package authProjectService

import "errors"

var (
	ErrUserRoleNotFound = errors.New("failed to find user role by this project")
)
