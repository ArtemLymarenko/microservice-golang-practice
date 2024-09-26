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

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now()
	u.SetUpdatedAt(time.Now())
}

func (u *User) SetUserInfo(firstName, lastName string) {
	u.UserInfo.SetFirstName(firstName)
	u.UserInfo.SetLastName(lastName)
	u.UserInfo.SetCreatedAt()
}
