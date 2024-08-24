package server

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Authenticate(token string) error {
	return nil
}
