package model

type Role string

const (
	SuperAdmin Role = "super_admin"
	Admin      Role = "admin"
	Dev        Role = "dev"
	Guest      Role = "guest"
)

type UserRole struct {
	UserId string
	Role   Role
}

func NewUserRole(userId string, role Role) UserRole {
	return UserRole{userId, role}
}
