package jwt

func (j *Manager) VerifyAccessToken(token string) (Claims, error) {
    claims, err := j.verify(token, j.cfg.AccessSecret)
    if err != nil {
        return Claims{}, err
    }

    if claims.TokenType != AccessToken {
        return Claims{}, ErrInvalidTokenType
    }

    return *claims, nil
}
