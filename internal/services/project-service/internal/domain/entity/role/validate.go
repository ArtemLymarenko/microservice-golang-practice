package role

import "errors"

var ErrRoleNotValid = errors.New("provided role is not valid")

func Validate(role Role) error {
	switch role {
	case Owner, Member:
		return nil
	default:
		return ErrRoleNotValid
	}
}
