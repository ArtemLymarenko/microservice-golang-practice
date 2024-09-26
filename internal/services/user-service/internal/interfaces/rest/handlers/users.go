package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-management-system/internal/user-service/internal/domain/model"
	"project-management-system/internal/user-service/internal/interfaces/mapper"
	"project-management-system/internal/user-service/internal/interfaces/rest/dto"
)

type UserService interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	Save(ctx context.Context, user model.User) error
}

type UsersHandler struct {
	userService UserService
}

func NewUsersHandler(userServ UserService) *UsersHandler {
	return &UsersHandler{userServ}
}

func (u *UsersHandler) Register(c *gin.Context) {
	var registerDto dto.RegisterUser
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := mapper.FromRegisterUserDTOToModel(registerDto)
	ctx := c.Request.Context()
	err := u.userService.Save(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you was registered"})
}
