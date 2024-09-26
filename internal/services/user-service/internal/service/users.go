package service

import (
	"context"
	"github.com/google/uuid"
	"project-management-system/internal/user-service/internal/domain/model"
	"strings"
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
	ctxWithTimeout, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.usersRepo.FindById(ctxWithTimeout, id)
}

func (u *UsersService) Save(ctx context.Context, user model.User) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	if strings.Trim(user.Id, " ") == "" {
		id := uuid.New()
		user.SetId(id.String())
	}

	user.SetCreatedAt()
	user.UserInfo.SetCreatedAt()

	return u.usersRepo.Save(ctxWithTimeout, user)
}

func (u *UsersService) Register(ctx context.Context, user model.User) error {
	//ctxWithTimeout, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	//defer cancel()

	return nil
}
