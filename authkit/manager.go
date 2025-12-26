package jwt

type Manager interface {
    GenerateAccessToken(userID string) (string, error)
    GenerateRefreshToken(userID string) (string, error)

    VerifyAccessToken(token string) (*Claims, error)
    VerifyRefreshToken(token string) (*Claims, error)
}
