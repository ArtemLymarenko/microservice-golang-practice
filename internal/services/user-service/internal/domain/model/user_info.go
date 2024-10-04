package model

import "time"

type UserInfo struct {
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserInfo) SetFirstName(firstName string) {
	u.FirstName = firstName
}

func (u *UserInfo) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *UserInfo) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}

func (u *UserInfo) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}
