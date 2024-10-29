package projectUserRepoPostgres

import "errors"

var (
	ErrSaveMember              = errors.New("failed to save project member")
	ErrMembersNotFound         = errors.New("failed to find member of project")
	ErrMemberWithRoleNotFound  = errors.New("failed to find member with role of project")
	ErrMembersWithRoleNotFound = errors.New("failed to find members with role of project")
	ErrDeleteMember            = errors.New("failed to delete project member")
)
