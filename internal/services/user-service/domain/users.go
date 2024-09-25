package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UserInfo  UserInfo  `json:"userInfo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserInfo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUserInfo(firstName, lastName string) UserInfo {
	return UserInfo{firstName, lastName}
}

func NewUser(email, password, firstName, lastName string) (*User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password must not be empty")
	}

	id := uuid.New()

	user := User{
		Id:        id,
		Email:     email,
		Password:  password,
		UserInfo:  NewUserInfo(firstName, lastName),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}
