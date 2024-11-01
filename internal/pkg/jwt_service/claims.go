package jwtService

import "github.com/golang-jwt/jwt/v5"

const (
	ClaimKeySubject   string = "Subject"
	ClaimKeyExpiresAt string = "ExpiresAt"
	ClaimKeyIssuer    string = "Issuer"
	ClaimKeyIssuedAt  string = "IssuedAt"
	ClaimKeyRole      string = "Role"
)

type Claims struct {
	claims jwt.MapClaims
}

func (c *Claims) GetClaim(key string) interface{} {
	return c.claims[key]
}
