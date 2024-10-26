package projectsPostgres

import "errors"

var (
	ErrProjectNotFound = errors.New("project was not found")
	ErrSaveProject     = errors.New("failed to save project")
	ErrDeleteProject   = errors.New("failed to delete project")
	ErrUpdateProject   = errors.New("failed to update project")
)
