package jwt

import "time"

type Algorithm string

const (
	HS256 Algorithm = "HS256"
)

type Config struct {
	Algorithm      Algorithm     `json:"algorithm"`
	Secret         string        `json:"secret"`
	Issuer         string        `json:"issuer"`
	Audience       string        `json:"audience"`
	AccessTokenTTL time.Duration `json:"accessTokenTTL"`
}
