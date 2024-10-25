package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	jwtService "project-management-system/internal/pkg/jwt-service"
	"project-management-system/internal/project-service/internal/config"
	v1Handlers "project-management-system/internal/project-service/internal/interface/rest/v1/handlers"
	"project-management-system/internal/project-service/internal/interface/rest/v1/middleware"
)

type Storage interface {
	GetConnection() (*sql.DB, error)
}

func MustGetGinRouter(storage Storage, cfg *config.Config) *gin.Engine {
	const ApiV1 = "/api/v1"

	const (
		Projects = "/projects"
		GetById  = "/:id"
	)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlers, err := v1Handlers.New(storage, 5000, cfg)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}

	jwtServ := jwtService.New(cfg.JWT.Secret, cfg.App.CodeName)

	apiPrivateV1Routes := router.Group(ApiV1)
	apiPrivateV1Routes.Use(middleware.Auth(jwtServ))
	{
		projects := apiPrivateV1Routes.Group(Projects)
		{
			projects.GET(GetById, handlers.ProjectHandler.GetProjectById)
		}
	}

	return router
}
