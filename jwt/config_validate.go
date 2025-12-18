package jwt

import "errors"

func (c Config) Validate() error {
	if c.Algorithm == "" {
		return errors.New("jwt : algorithm is required")
	}
	if c.Algorithm == "HS256" && c.Secret == "" {
		return errors.New("jwt : secret is required")
	}
	if c.AccessTokenTTL <= 0 {
		return errors.New("jwt : access token TTL must be greater than zero")
	}
	return nil
}
