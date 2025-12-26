package jwt

type Manager struct {
    cfg Config
}

func NewJWTManager(cfg Config) *Manager {
    return &Manager{cfg}
}

func (j *Manager) GenerateAccessToken(userID string) (string, error) {
    return j.generateToken(userID, AccessToken, j.cfg.AccessTTL, j.cfg.AccessSecret)
}

func (j *Manager) GenerateRefreshToken(userID string) (string, error) {
    return j.generateToken(userID, RefreshToken, j.cfg.RefreshTokenTTL, j.cfg.RefreshSecret)
}

func (j *Manager) VerifyRefreshToken(token string) (string, error) {
    claims, err := j.verify(token, j.cfg.RefreshSecret)
    if err != nil {
        return "", err
    }

    if claims.TokenType != RefreshToken {
        return "", ErrInvalidTokenType
    }

    return claims.UserID, nil
}
