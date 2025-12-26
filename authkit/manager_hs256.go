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

func (m *hs256Manager) GenerateAccessToken(userID string) (string, error) {
    return m.generate(userID, AccessToken, m.config.AccessTokenTTL, m.config.AccessSecret)
}

func (m *hs256Manager) GenerateRefreshToken(userID string) (string, error) {
    return m.generate(userID, RefreshToken, m.config.RefreshTokenTTL, m.config.RefreshSecret)
}

func (m *hs256Manager) generate(
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
            Issuer:    m.config.Issuer,
            Audience:  jwt.ClaimStrings{m.config.Audience},
            IssuedAt:  jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func (m *hs256Manager) VerifyAccessToken(tokenStr string) (*Claims, error) {
    claims, err := m.verify(tokenStr, m.config.AccessSecret)
    if err != nil {
        return nil, err
    }

    if claims.TokenType != AccessToken {
        return nil, ErrInvalidTokenType
    }

    return claims, nil
}

func (m *hs256Manager) VerifyRefreshToken(tokenStr string) (*Claims, error) {
    claims, err := m.verify(tokenStr, m.config.RefreshSecret)
    if err != nil {
        return nil, err
    }

    if claims.TokenType != RefreshToken {
        return nil, ErrInvalidTokenType
    }

    return claims, nil
}

func (m *hs256Manager) verify(tokenStr string, secret string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(
        tokenStr,
        claims,
        func(token *jwt.Token) (interface{}, error) {
            if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
                return nil, errors.New("jwt: unexpected signing method")
            }
            return []byte(secret), nil
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
