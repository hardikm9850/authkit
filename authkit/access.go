package jwt

func (j *JWTManager) VerifyAccessToken(token string) (string, error) {
    claims, err := j.verify(token, j.cfg.AccessSecret)
    if err != nil {
        return "", err
    }

    if claims.TokenType != AccessToken {
        return "", ErrInvalidTokenType
    }

    return claims.UserID, nil
}
