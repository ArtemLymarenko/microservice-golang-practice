package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	jwtService "project-management-system/internal/pkg/jwt_service"
	sqlStorage "project-management-system/internal/pkg/sql_storage"
	"project-management-system/internal/project-service/internal/config"
	projectUserRepoPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/project_user"
	projectsRepoPostgres "project-management-system/internal/project-service/internal/infrastructure/repository/postgres/projects"
	v1Handlers "project-management-system/internal/project-service/internal/interface/rest/v1/handlers"
	"project-management-system/internal/project-service/internal/interface/rest/v1/middleware"
	projectsService "project-management-system/internal/project-service/internal/service/projects"
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

	//sql
	sqlManager := sqlStorage.NewSqlDB(connection)
	txManager := sqlStorage.NewTxManager(connection)

	//repos
	projectRepo := projectsRepoPostgres.New(sqlManager)
	projectUserRepo := projectUserRepoPostgres.New(sqlManager)

	//third-party
	jwtServ := jwtService.New(cfg.JWT.Secret, cfg.App.CodeName) //move logic to api gateway
	validatorService := validator.New()

	//services
	projectServ := projectsService.New(projectRepo, projectUserRepo, txManager, validatorService)

	//handlers
	handlers := v1Handlers.New(projectServ)

	apiPrivateV1Routes := router.Group(ApiV1)
	apiPrivateV1Routes.Use(middleware.Auth(jwtServ))
	{
		projects := apiPrivateV1Routes.Group(Projects)
		projects.Use(middleware.RetrieveProjectRoleByUser(projectUserRepo))
		{
			projects.GET(GetById, handlers.ProjectHandler.GetProjectById)
		}
	}

	return router
}
