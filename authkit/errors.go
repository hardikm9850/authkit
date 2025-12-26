package jwt

import "errors"

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrInvalidTokenType = errors.New("invalid token type")
	ErrExpiredToken     = errors.New("jwt: token expired")
	ErrInvalidIssuer    = errors.New("jwt: invalid issuer")
	ErrInvalidAudience  = errors.New("jwt: invalid audience")
)
