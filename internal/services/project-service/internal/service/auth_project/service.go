package authProjectService

import (
	"context"
	jwtService "project-management-system/internal/pkg/jwt_service"
	"project-management-system/internal/project-service/internal/domain/entity/project"
	"project-management-system/internal/project-service/internal/domain/entity/user"
	"project-management-system/internal/project-service/internal/domain/repository/persistent"
	"project-management-system/internal/project-service/internal/interface/rest/dto"
	"project-management-system/internal/project-service/internal/service"
	"time"
)

type JwtService interface {
	Generate(subject string, exp time.Duration, additionalFields map[string]interface{}) (string, error)
}

type AuthProjectService struct {
	projectUserRepo persistent.ProjectUserRepository
	jwtService      JwtService
}

func New(projectUserRepo persistent.ProjectUserRepository, jwtService JwtService) *AuthProjectService {
	return &AuthProjectService{
		projectUserRepo: projectUserRepo,
		jwtService:      jwtService,
	}
}

func (a *AuthProjectService) IssueProjectToken(
	ctx context.Context,
	userId user.Id,
	projectId project.Id,
) (result dto.AuthProjectResponseDto, err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, service.TIMEOUT)
	defer cancel()

	role, err := a.projectUserRepo.FindUserRoleByProject(ctxWithTimeout, userId, projectId)
	if err != nil {
		return result, ErrUserRoleNotFound
	}

	subject := string(projectId)
	duration := 4 * time.Hour
	additionalClaims := map[string]interface{}{
		jwtService.ClaimKeyRole: role,
	}
	token, err := a.jwtService.Generate(subject, duration, additionalClaims)
	if err != nil {
		return result, err
	}

	return dto.AuthProjectResponseDto{
		ProjectToken: token,
		ExpiresIn:    duration.String(),
	}, nil

}
