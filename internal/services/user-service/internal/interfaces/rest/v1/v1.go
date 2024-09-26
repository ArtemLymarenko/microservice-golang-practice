package v1

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter(handlers *Handlers) *gin.Engine {
	const (
		Users = "/users"
	)

	router := gin.Default()

	router.POST(Users, handlers.UsersHandler.Register)

	return router
}
