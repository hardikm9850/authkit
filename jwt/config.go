package jwt

import "time"

type Algorithm string

const (
	HS256 Algorithm = "HS256"
)

type Config struct {
	Algorithm       Algorithm `json:"algorithm"`
	AccessSecret    string    `json:"accessSecret"`
	RefreshSecret   string    `json:"refreshSecret"`
	AccessTTL       time.Duration
	RefreshTokenTTL time.Duration
	Issuer          string        `json:"issuer"`
	Audience        string        `json:"audience"`
	AccessTokenTTL  time.Duration `json:"accessTokenTTL"`
}
