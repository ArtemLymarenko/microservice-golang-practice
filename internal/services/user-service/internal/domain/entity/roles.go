package entity

type RolePermission string

const (
	PRIMARY_ADMIN RolePermission = "primary_admin"
	ADMIN         RolePermission = "admin"
	GUEST         RolePermission = "guest"
)

type Role struct {
	Id             string         `json:"id"`
	RolePermission RolePermission `json:"rolePermission"`
}
