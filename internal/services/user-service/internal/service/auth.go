package service

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"project-management-system/internal/user-service/internal/config"
	"project-management-system/internal/user-service/internal/domain/model"
	"project-management-system/internal/user-service/internal/interfaces/rest/dto"
	jwtService "project-management-system/internal/user-service/pkg/jwt-service"
	"time"
)

type UsersServ interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Save(ctx context.Context, user model.User) error
}

type JWTServ interface {
	Generate(userId string, expirationTime time.Duration) (string, error)
	GenerateTokenAsync(userId string, exp time.Duration) chan string
	Verify(tokenToCheck string) (*jwtService.Claims, error)
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
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	id := uuid.New().String()
	user.SetId(id)

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost+bcrypt.MinCost)
	if err != nil {
		return nil, ErrHashingPassword
	}
	user.SetPassword(string(password))

	err = a.usersService.Save(ctxTimeout, user)
	if err != nil {
		return nil, err
	}

	tokens, err := a.generateTokens(id)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a *AuthService) Login(ctx context.Context, user model.User) (*dto.AuthResponse, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	foundUser, err := a.usersService.FindByEmail(ctxTimeout, user.Email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	hashedPassword := []byte(foundUser.Password)
	userPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword, userPassword)
	if err != nil {
		return nil, ErrPasswordsNotMatch
	}

	tokens, err := a.generateTokens(foundUser.Id)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a *AuthService) IssueTokens(ctx context.Context, refreshToken string) (*dto.AuthResponse, error) {
	claims, err := a.jwtService.Verify(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := a.usersService.FindById(ctx, claims.Subject)
	if err != nil {
		return nil, err
	}

	tokens, err := a.generateTokens(user.Id)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a *AuthService) generateTokens(
	sub string,
) (*dto.AuthResponse, error) {
	accessChan := a.jwtService.GenerateTokenAsync(sub, a.jwtConfig.AccessExp)
	refreshChan := a.jwtService.GenerateTokenAsync(sub, a.jwtConfig.RefreshExp)
	defer close(accessChan)
	defer close(refreshChan)

	accessToken, refreshToken := <-accessChan, <-refreshChan
	if accessToken == "" || refreshToken == "" {
		return nil, ErrGenerateTokens
	}

	return &dto.AuthResponse{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  a.jwtConfig.AccessExp.String(),
		RefreshExpiresIn: a.jwtConfig.RefreshExp.String(),
	}, nil
}
