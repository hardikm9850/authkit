package jwt

import "errors"

func (c Config) Validate() error {
	if c.Algorithm == "" {
		return errors.New("jwt : algorithm is required")
	}
	if c.Algorithm == "HS256" && c.AccessSecret == "" {
		return errors.New("jwt : access secret is required")
	}
	if c.AccessTokenTTL <= 0 {
		return errors.New("jwt : access token TTL must be greater than zero")
	}
	if c.RefreshSecret == "" {
		return errors.New("jwt : refresh secret is required")
	}
	if c.RefreshTokenTTL <= 0 {
		return errors.New("jwt : refresh token TTL must be greater than zero")
	}
	return nil
}
