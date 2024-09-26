package jwtService

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

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

	return token.SignedString([]byte(jwts.secret))
}

func (jwts JWTService) Verify(tokenToCheck string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenToCheck, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwts.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
