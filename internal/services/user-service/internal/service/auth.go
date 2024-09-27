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
	FindByEmail(ctx context.Context, email string) (*model.User, error)
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

func (a *AuthService) Register(ctx context.Context, user model.User) (*dto.AuthResponse, error) {
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
		return nil, errors.New("failed to hash password")
	}
	user.SetPassword(string(password))

	err = a.usersService.Save(ctxWithTimeout, user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  a.jwtConfig.AccessExp.String(),
		RefreshExpiresIn: a.jwtConfig.RefreshExp.String(),
	}, nil
}

func (a *AuthService) Login(ctx context.Context, user model.User) (*dto.AuthResponse, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	foundUser, err := a.usersService.FindByEmail(ctxWithTimeout, user.Email)
	if err != nil {
		return nil, errors.New("user was not found")
	}

	hashedPassword := []byte(foundUser.Password)
	userPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword, userPassword)
	if err != nil {
		return nil, errors.New("passwords do not match")
	}

	accessToken, refreshToken, err := a.generateTokens(foundUser.Id, a.jwtConfig.AccessExp, a.jwtConfig.RefreshExp)
	return &dto.AuthResponse{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  a.jwtConfig.AccessExp.String(),
		RefreshExpiresIn: a.jwtConfig.RefreshExp.String(),
	}, nil
}

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
