package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (j *Manager) generateToken(
	userID string,
	tokenType TokenType,
	ttl time.Duration,
	secret string,
) (string, error) {

	now := time.Now()

	claims := Claims{
		UserID:    userID,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
