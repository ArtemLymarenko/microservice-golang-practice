package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/domain/repository/postgres"
	"project-management-system/internal/user-service/internal/service"
	"time"
)

type Storage interface {
	GetConnection() (*sql.DB, error)
}

type UsersHandler interface {
	Register(c *gin.Context)
}

type Handlers struct {
	UsersHandler UsersHandler
}

func New(storage Storage, serviceTimeout time.Duration) (*Handlers, error) {
	connection, err := storage.GetConnection()
	if err != nil {
		return nil, err
	}

	//repos
	userInfoRepo := postgres.NewUserInfoRepository(connection)
	usersRepo := postgres.NewUsersRepository(connection, userInfoRepo)

	//services
	userService := service.NewUsersService(usersRepo, serviceTimeout)

	return &Handlers{
		UsersHandler: NewUsersHandler(userService),
	}, err
}
