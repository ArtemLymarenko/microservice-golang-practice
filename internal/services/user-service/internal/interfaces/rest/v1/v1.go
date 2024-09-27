package v1

import (
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/interfaces/rest/v1/handlers"
)

func InitializeRouter(handlers *handlers.Handlers) *gin.Engine {
	const (
		AuthRegister = "/register"
		AuthLogin    = "/login"
	)

	router := gin.Default()

	router.POST(AuthRegister, handlers.AuthHandler.Register)
	router.POST(AuthLogin, handlers.AuthHandler.Login)

	return router
}
