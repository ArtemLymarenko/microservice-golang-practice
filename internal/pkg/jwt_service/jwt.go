package jwtService

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	*jwt.RegisteredClaims
}

type JWTService struct {
	secret string
	issuer string
}

func New(secret, issuer string) *JWTService {
	return &JWTService{secret, issuer}
}

func (jwtService *JWTService) Generate(userId string, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		Issuer:    jwtService.issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	signedToken, err := token.SignedString([]byte(jwtService.secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (jwtService *JWTService) GenerateTokenAsync(
	userId string,
	exp time.Duration,
) chan string {
	tokenChan := make(chan string)
	go func() {
		token, err := jwtService.Generate(userId, exp)
		if err != nil {
			tokenChan <- ""
			return
		}

		tokenChan <- token
	}()

	return tokenChan
}

func (jwtService *JWTService) Verify(token string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtService.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok && parsedToken.Valid {
		return &Claims{claims}, nil
	}

	return nil, ErrInvalidToken
}
