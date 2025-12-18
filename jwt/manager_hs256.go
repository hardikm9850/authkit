package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type hs256Manager struct {
	config Config
}

func NewManager(config Config) (Manager, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &hs256Manager{config: config}, nil
}

func (m *hs256Manager) Generate(claims Claims) (string, error) {
	now := time.Now()

	claims.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    m.config.Issuer,
		Audience:  jwt.ClaimStrings{m.config.Audience},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.config.AccessTokenTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.config.Secret))
}

func (m *hs256Manager) Verify(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != string(HS256) {
				return nil, errors.New("jwt: unexpected signing method")
			}
			return []byte(m.config.Secret), nil
		},
		jwt.WithAudience(m.config.Audience),
		jwt.WithIssuer(m.config.Issuer),
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("jwt: invalid token")
	}

	return claims, nil
}
