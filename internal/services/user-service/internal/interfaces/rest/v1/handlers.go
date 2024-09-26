package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/domain/repository/postgres"
	"project-management-system/internal/user-service/internal/interfaces/rest/handlers"
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

func InitializeHandlers(storage Storage, serviceTimeout time.Duration) (*Handlers, error) {
	connection, err := storage.GetConnection()
	if err != nil {
		return nil, err
	}

	//repos
	userInfoRepo := postgres.NewUserInfoRepository(connection)
	usersRepo := postgres.NewUsersRepository(connection, userInfoRepo)

	//services
	userService := service.NewUsersService(usersRepo, serviceTimeout)

	//handlers
	usersHandler := handlers.NewUsersHandler(userService)

	return &Handlers{
		usersHandler,
	}, err
}
