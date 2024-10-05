package model

import (
	"time"
)

type User struct {
	Id        string
	Email     string
	Password  string
	UserInfo  UserInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) SetId(id string) {
	u.Id = id
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}
