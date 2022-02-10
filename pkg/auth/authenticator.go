package auth

import (
	"github.com/keptn/cli2/pkg/store"
)

type AuthOptions struct {
	Endpoint        string
	APIToken        string
	SSO             bool
	SSOLogout       bool
	SSODiscovery    string
	SSOClientID     string
	SSOScopes       []string
	SSOClientSecret string
}

type Authenticator interface {
	Auth(AuthOptions) error
}

type TokenAuthenticator struct {
	authStore store.AuthStore
}

func New(store store.AuthStore) *TokenAuthenticator {
	return &TokenAuthenticator{authStore: store}
}

func (a *TokenAuthenticator) Auth(opts AuthOptions) error {
	//TODO implement me
	panic("implement me")
}
