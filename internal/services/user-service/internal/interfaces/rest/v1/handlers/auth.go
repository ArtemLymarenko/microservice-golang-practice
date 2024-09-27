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
	Register(ctx context.Context, user model.User) (*dto.AuthResponse, error)
	Login(ctx context.Context, user model.User) (*dto.AuthResponse, error)
}

type AuthHandlerImpl struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandlerImpl {
	return &AuthHandlerImpl{authService}
}

func (a *AuthHandlerImpl) Register(c *gin.Context) {
	var registerDto dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := mapper.FromRegisterUserDTOToModel(registerDto)
	ctx := c.Request.Context()
	authResponse, err := a.authService.Register(ctx, user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authResponse)
}

func (a *AuthHandlerImpl) Login(c *gin.Context) {
	var loginDto dto.LoginUserRequest
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := mapper.FromLoginUserDTOToModel(loginDto)
	ctx := c.Request.Context()
	authResponse, err := a.authService.Login(ctx, user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authResponse)
}
