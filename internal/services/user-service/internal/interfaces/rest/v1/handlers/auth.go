package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-management-system/internal/user-service/internal/domain/model"
	"project-management-system/internal/user-service/internal/interfaces/mapper"
	"project-management-system/internal/user-service/internal/interfaces/rest/dto"
)

type AuthService interface {
	Register(ctx context.Context, user model.User) (*dto.RegisterUserResponse, error)
}

type authHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *authHandler {
	return &authHandler{authService}
}

func (a *authHandler) Register(c *gin.Context) {
	var registerDto dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := mapper.FromRegisterUserDTOToModel(registerDto)
	ctx := c.Request.Context()
	registerUserResponse, err := a.authService.Register(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registerUserResponse)
}
