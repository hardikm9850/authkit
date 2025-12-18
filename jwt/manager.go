package jwt

type Manager interface {
	Generate(claims Claims) (string, error)
	Verify(token string) (*Claims, error)
}
