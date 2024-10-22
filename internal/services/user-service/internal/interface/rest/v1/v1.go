package v1

import (
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/interface/rest/v1/handlers"
)

func GetGinRouter(handlers *handlers.Handlers) *gin.Engine {
	const ApiV1 = "/api/v1"

	const (
		IssueTokens  = "/issue-tokens"
		AuthRegister = "/register"
		AuthLogin    = "/login"
	)

	router := gin.Default()
	apiV1Routes := router.Group(ApiV1)
	{
		apiV1Routes.POST(AuthRegister, handlers.AuthHandler.Register)
		apiV1Routes.POST(AuthLogin, handlers.AuthHandler.Login)
		apiV1Routes.POST(IssueTokens, handlers.AuthHandler.IssueTokens)
	}

	return router
}
