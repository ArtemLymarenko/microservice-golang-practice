package model

import (
	"github.com/google/uuid"
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

func NewUser(email, password, firstName, lastName string) (User, error) {
	id := uuid.New()
	return User{
		Id:        id.String(),
		Email:     email,
		Password:  password,
		UserInfo:  NewUserInfo(firstName, lastName),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
