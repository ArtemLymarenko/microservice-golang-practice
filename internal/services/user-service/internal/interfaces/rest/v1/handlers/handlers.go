package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/config"
	"project-management-system/internal/user-service/internal/domain/repository/postgres"
	"project-management-system/internal/user-service/internal/service"
	jwtService "project-management-system/internal/user-service/pkg/jwt-service"
	"time"
)

type Storage interface {
	GetConnection() (*sql.DB, error)
}

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type Handlers struct {
	AuthHandler AuthHandler
}

func New(storage Storage, serviceTimeout time.Duration, cfg *config.Config) (*Handlers, error) {
	connection, err := storage.GetConnection()
	if err != nil {
		return nil, err
	}

	//repos
	userInfoRepo := postgres.NewUserInfoRepository(connection)
	usersRepo := postgres.NewUsersRepository(connection, userInfoRepo)

	//third-party
	jwtServ := jwtService.New(cfg.JWT.Secret, cfg.App.CodeName)

	//services
	userService := service.NewUsersService(usersRepo, serviceTimeout)
	authService := service.NewAuthService(cfg.JWT, userService, jwtServ, serviceTimeout)

	return &Handlers{
		AuthHandler: NewAuthHandler(authService),
	}, err
}
