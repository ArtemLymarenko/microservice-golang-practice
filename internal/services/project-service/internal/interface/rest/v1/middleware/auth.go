package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	jwtService "project-management-system/internal/pkg/jwt-service"
	"project-management-system/internal/project-service/internal/interface/rest/ctxkey"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
	"strings"
)

var (
	ErrFailedToGetToken      = errors.New("failed to get token")
	ErrFailedToAuthorizeUser = errors.New("failed to authorize user")
)

type JWTService interface {
	Verify(token string) (*jwtService.Claims, error)
}

func Auth(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		bearer, token, found := strings.Cut(authHeader, " ")
		if !found && bearer != "Bearer" || token == "" {
			c.JSON(http.StatusUnauthorized, dto.NewResponseErr(ErrFailedToGetToken))
			c.Abort()
			return
		}

		claims, err := jwtService.Verify(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.NewResponseErr(ErrFailedToAuthorizeUser))
			c.Abort()
			return
		}

		c.Set(string(ctxkey.UserId), claims.Subject)
		c.Next()
	}
}
