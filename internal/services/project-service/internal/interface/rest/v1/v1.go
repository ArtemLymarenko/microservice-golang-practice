package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	jwtService "project-management-system/internal/pkg/jwt-service"
	"project-management-system/internal/project-service/internal/config"
	"project-management-system/internal/project-service/internal/infrastructure/repository/postgres"
	projectUserRepoPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/project_user"
	projectsRepoPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/projects"
	v1Handlers "project-management-system/internal/project-service/internal/interface/rest/v1/handlers"
	"project-management-system/internal/project-service/internal/interface/rest/v1/middleware"
	projectService "project-management-system/internal/project-service/internal/service/projects"
)

func MustGetGinRouter(connection *sql.DB, cfg *config.Config) *gin.Engine {
	const ApiV1 = "/api/v1"

	const (
		Projects = "/projects"
		GetById  = "/:id"
	)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtServ := jwtService.New(cfg.JWT.Secret, cfg.App.CodeName)

	//repos
	txManager := postgres.NewTxManager(connection)
	projectRepo := projectsRepoPostgres.New(connection)
	projectUserRepo := projectUserRepoPostgres.New(connection, txManager, projectRepo)

	//services
	projectServ := projectService.New(projectRepo, projectUserRepo, txManager)

	//handlers
	handlers := v1Handlers.New(projectServ)

	apiPrivateV1Routes := router.Group(ApiV1)
	apiPrivateV1Routes.Use(middleware.Auth(jwtServ))
	{
		projects := apiPrivateV1Routes.Group(Projects)
		apiPrivateV1Routes.Use(middleware.RetrieveProjectRoleByUser(projectUserRepo))
		{
			projects.GET(GetById, handlers.ProjectHandler.GetProjectById)
		}
	}

	return router
}
