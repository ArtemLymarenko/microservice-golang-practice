package model

import (
	"errors"
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

func NewUser(email, password, firstName, lastName string) (*User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password must not be empty")
	}

	id := uuid.New()

	user := User{
		Id:        id.String(),
		Email:     email,
		Password:  password,
		UserInfo:  NewUserInfo(firstName, lastName),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}
