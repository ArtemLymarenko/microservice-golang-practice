package v1

import (
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/interfaces/rest/v1/handlers"
)

func InitializeRouter(handlers *handlers.Handlers) *gin.Engine {
	const (
		IssueTokens  = "/issue-tokens"
		AuthRegister = "/register"
		AuthLogin    = "/login"
	)

	router := gin.Default()
	apiGroup := router.Group("/api/v1")

	publicRoutes := apiGroup.Group("/")

	publicRoutes.POST(AuthRegister, handlers.AuthHandler.Register)
	publicRoutes.POST(AuthLogin, handlers.AuthHandler.Login)
	publicRoutes.POST(IssueTokens, handlers.AuthHandler.IssueTokens)

	return router
}
