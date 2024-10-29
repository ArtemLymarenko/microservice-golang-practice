package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/role"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/interface/rest/ctxkey"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
	"time"
)

var (
	ErrFailedToGetUserData = errors.New("failed to get user data")
)

type ProjectUserRepository interface {
	FindUserRoleByProject(
		ctx context.Context,
		userId user.Id,
		projectId project.Id,
	) (result role.Role, err error)
}

func RetrieveProjectRoleByUser(projectUserRepo ProjectUserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := c.Param("id")
		userId, exists := c.Get(ctxkey.UserId)
		if !exists || projectId == "" {
			c.JSON(http.StatusUnauthorized, dto.NewResponseErr(ErrFailedToGetUserData))
			c.Abort()
			return
		}

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		userProjectRole, err := projectUserRepo.FindUserRoleByProject(
			ctxWithTimeout,
			userId.(user.Id),
			project.Id(projectId),
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.NewResponseErr(ErrFailedToGetUserData))
			c.Abort()
			return
		}

		c.Set(ctxkey.ProjectRole, userProjectRole)
		c.Next()
	}
}
