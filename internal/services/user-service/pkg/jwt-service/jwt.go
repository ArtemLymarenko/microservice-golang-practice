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

func New(secret string, issuer string) JWTService {
	return JWTService{secret, issuer}
}

func (jwts JWTService) Generate(userId string, expirationTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		Issuer:    jwts.issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	signedToken, err := token.SignedString([]byte(jwts.secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (jwts JWTService) GenerateTokenAsync(
	userId string,
	exp time.Duration,
	tokenChan chan string,
) {
	token, err := jwts.Generate(userId, exp)
	if err != nil {
		return
	}

	tokenChan <- token
}

func (jwts JWTService) Verify(token string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwts.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok && parsedToken.Valid {
		return &Claims{claims}, nil
	}

	return nil, ErrInvalidToken
}
