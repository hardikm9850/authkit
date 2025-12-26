package jwt

import (
    "github.com/golang-jwt/jwt/v5"
)

func (j *Manager) verify(tokenString string, secret string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(
        tokenString,
        &Claims{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        },
    )

    if err != nil || !token.Valid {
        return nil, ErrInvalidToken
    }

    claims, ok := token.Claims.(*Claims)
    if !ok {
        return nil, ErrInvalidToken
    }

    return claims, nil
}
