package v1

import (
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/interface/rest/v1/handlers"
)

const (
	Root  = "/"
	ApiV1 = "/api/v1"
)

func InitializeRouter(handlers *handlers.Handlers) *gin.Engine {
	const (
		IssueTokens  = "/issue-tokens"
		AuthRegister = "/register"
		AuthLogin    = "/login"
	)

	router := gin.Default()
	apiV1Routes := router.Group(ApiV1)

	publicRoutes := apiV1Routes.Group(Root)
	publicRoutes.POST(AuthRegister, handlers.AuthHandler.Register)
	publicRoutes.POST(AuthLogin, handlers.AuthHandler.Login)
	publicRoutes.POST(IssueTokens, handlers.AuthHandler.IssueTokens)

	return router
}
