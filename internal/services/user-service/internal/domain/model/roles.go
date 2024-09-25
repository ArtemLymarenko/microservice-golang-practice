package model

import (
	"github.com/google/uuid"
)

type RoleName string

const (
	SUPER_ADMIN RoleName = "super_admin"
	ADMIN       RoleName = "admin"
	DEV         RoleName = "dev"
	GUEST       RoleName = "guest"
)

type Role struct {
	Id       uuid.UUID
	RoleName RoleName
}

func NewRole(roleName RoleName) Role {
	id := uuid.New()
	return Role{id, roleName}
}
