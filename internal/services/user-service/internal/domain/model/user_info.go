package model

import "time"

type UserInfo struct {
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserInfo(firstName, lastName string) UserInfo {
	return UserInfo{
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
