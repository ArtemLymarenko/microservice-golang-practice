package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"project-management-system/internal/user-service/internal/config"
	"project-management-system/internal/user-service/internal/domain/model"
	"project-management-system/internal/user-service/internal/interfaces/rest/dto"
	"time"
)

type UsersServ interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	Save(ctx context.Context, user model.User) error
}

type JWTServ interface {
	Generate(userId string, expirationTime time.Duration) (string, error)
	GenerateTokenAsync(userId string, exp time.Duration, tokenChan chan string)
	Verify(tokenToCheck string) (*jwt.RegisteredClaims, error)
}

type AuthService struct {
	jwtConfig    config.JWT
	usersService UsersServ
	jwtService   JWTServ
	ctxTimeout   time.Duration
}

func NewAuthService(
	jwtConfig config.JWT,
	usersService UsersServ,
	jwtService JWTServ,
	ctxTimeout time.Duration,
) *AuthService {
	return &AuthService{
		jwtConfig:    jwtConfig,
		usersService: usersService,
		jwtService:   jwtService,
		ctxTimeout:   ctxTimeout,
	}
}

func (a *AuthService) Register(ctx context.Context, user model.User) (*dto.RegisterUserResponse, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	id := uuid.New().String()
	user.SetId(id)

	accessToken, refreshToken, err := a.generateTokens(id, a.jwtConfig.AccessExp, a.jwtConfig.RefreshExp)
	if err != nil {
		return nil, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost+bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	user.SetPassword(string(password))

	err = a.usersService.Save(ctxWithTimeout, user)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  a.jwtConfig.AccessExp.String(),
		RefreshExpiresIn: a.jwtConfig.RefreshExp.String(),
	}, nil
}

func (a *AuthService) Login() {}

func (a *AuthService) IssueTokens(refreshToken string) {}

func (a *AuthService) generateTokens(
	userId string,
	accessExp, refreshExp time.Duration,
) (accessToken string, refreshToken string, err error) {
	accessChan := make(chan string)
	refreshChan := make(chan string)
	defer close(accessChan)
	defer close(refreshChan)

	go a.jwtService.GenerateTokenAsync(userId, accessExp, accessChan)
	go a.jwtService.GenerateTokenAsync(userId, refreshExp, refreshChan)

	accessToken, refreshToken = <-accessChan, <-refreshChan
	if accessToken == "" || refreshToken == "" {
		return accessToken, refreshToken, errors.New("failed to generate tokens")
	}

	return accessToken, refreshToken, nil
}
