package v1

import (
	"github.com/gin-gonic/gin"
	"project-management-system/internal/user-service/internal/interfaces/rest/v1/handlers"
)

func InitializeRouter(handlers *handlers.Handlers) *gin.Engine {
	const (
		Users = "/users"
	)

	router := gin.Default()

	router.POST(Users, handlers.UsersHandler.Register)

	return router
}
