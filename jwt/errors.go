package jwt

import "errors"

var (
	ErrInvalidToken    = errors.New("jwt: invalid token")
	ErrExpiredToken    = errors.New("jwt: token expired")
	ErrInvalidIssuer   = errors.New("jwt: invalid issuer")
	ErrInvalidAudience = errors.New("jwt: invalid audience")
)
