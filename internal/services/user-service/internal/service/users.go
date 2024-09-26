package service

import (
	"context"
	"project-management-system/internal/user-service/internal/domain/model"
	"time"
)

type UsersRepo interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	Save(ctx context.Context, user model.User) error
}

type UsersService struct {
	usersRepo  UsersRepo
	ctxTimeout time.Duration
}

func NewUsersService(usersRepo UsersRepo, ctxTimeout time.Duration) *UsersService {
	return &UsersService{usersRepo, ctxTimeout}
}

func (u *UsersService) FindById(ctx context.Context, id string) (*model.User, error) {
	timeout, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.usersRepo.FindById(timeout, id)
}

func (u *UsersService) Save(ctx context.Context, user model.User) error {
	timeout, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.usersRepo.Save(timeout, user)
}
