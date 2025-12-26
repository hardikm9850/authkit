package jwt

type JWTManager struct {
    cfg Config
}

func NewJWTManager(cfg Config) *JWTManager {
	return &JWTManager{cfg}
}

func (j *JWTManager) GenerateAccessToken(userID string) (string, error) {
    return j.generateToken(userID, AccessToken, j.cfg.AccessTTL, j.cfg.AccessSecret)
}

func (j *JWTManager) GenerateRefreshToken(userID string) (string, error) {
    return j.generateToken(userID, RefreshToken, j.cfg.RefreshTokenTTL, j.cfg.RefreshSecret)
}

func (j *JWTManager) VerifyRefreshToken(token string) (string, error) {
    claims, err := j.verify(token, j.cfg.RefreshSecret)
    if err != nil {
        return "", err
    }

    if claims.TokenType != RefreshToken {
        return "", ErrInvalidTokenType
    }

    return claims.UserID, nil
}
