package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"project-management-system/internal/user-service/internal/domain/model"
	"time"
)

type UsersServ interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	Save(ctx context.Context, user model.User) error
}

type JWTServ interface {
	Generate(userId string, expirationTime time.Duration) (string, error)
	Verify(tokenToCheck string) (*jwt.RegisteredClaims, error)
}

type AuthService struct {
	usersService UsersServ
	jwtService   JWTServ
}

func NewAuthService(usersService UsersServ, jwtService JWTServ) *AuthService {
	return &AuthService{usersService, jwtService}
}

func (a *AuthService) Register() {}

func (a *AuthService) Login() {}

func (a *AuthService) IssueTokens(refreshToken string) {}

func (a *AuthService) generateTokenAsync(
	userId string,
	exp time.Duration,
	tokenChan chan string,
) {
	token, err := a.jwtService.Generate(userId, exp)
	if err != nil {
		tokenChan <- ""
		return
	}

	tokenChan <- token
}

func (a *AuthService) generateTokens(userId string, accessExp, refreshExp time.Duration) (string, string, error) {
	accessChan := make(chan string)
	refreshChan := make(chan string)
	defer close(accessChan)
	defer close(refreshChan)

	go a.generateTokenAsync(userId, accessExp, accessChan)
	go a.generateTokenAsync(userId, refreshExp, refreshChan)

	accessToken, refreshToken := <-accessChan, <-refreshChan
	if accessToken == "" || refreshToken == "" {
		return "", "", errors.New("failed to generate tokens")
	}

	return accessToken, refreshToken, nil
}
