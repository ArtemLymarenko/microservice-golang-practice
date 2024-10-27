package projectUserRepoPostgres

import "errors"

var (
	ErrSaveProjectWithUser = errors.New("failed to save project with user")
	ErrSaveMember          = errors.New("failed to save project member")
)
