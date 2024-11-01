package jwtService

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	secret string
	issuer string
}

func New(secret, issuer string) *JWTService {
	return &JWTService{secret, issuer}
}

func (jwtService *JWTService) addCustomFields(
	claims jwt.MapClaims,
	customFields map[string]interface{},
) jwt.MapClaims {
	for key, value := range customFields {
		claims[key] = value
	}
	return claims
}

func (jwtService *JWTService) Generate(
	subject string,
	exp time.Duration,
	additionalFields map[string]interface{},
) (string, error) {
	claims := jwt.MapClaims{
		"Subject":   subject,
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(exp)),
		"Issuer":    jwtService.issuer,
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
	}

	claims = jwtService.addCustomFields(claims, additionalFields)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtService.secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (jwtService *JWTService) GenerateTokenAsync(
	userId string,
	exp time.Duration,
	additionalFields map[string]interface{},
) chan string {
	tokenChan := make(chan string)
	go func() {
		token, err := jwtService.Generate(userId, exp, additionalFields)
		if err != nil {
			tokenChan <- ""
			return
		}

		tokenChan <- token
	}()

	return tokenChan
}

func (jwtService *JWTService) Verify(token string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtService.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return &Claims{claims}, nil
	}

	return nil, ErrInvalidToken
}
