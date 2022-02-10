package auth

type TestAuthenticator struct {
	AuthenticatorFunc func(options AuthOptions) error
}

func (a *TestAuthenticator) Auth(options AuthOptions) error {
	return a.AuthenticatorFunc(options)
}
