package jwt

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID    string    `json:"userId"`
	Roles     []string  `json:"roles"`
	TokenType TokenType `json:"typ"`
	jwt.RegisteredClaims
}
