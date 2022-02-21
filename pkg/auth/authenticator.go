package auth

import (
	api "github.com/keptn/go-utils/pkg/api/utils"
)

type Authenticator interface {
	Auth(apiEndpoint, apiToken string) (api.KeptnInterface, error)
	ReAuth() (api.KeptnInterface, error)
}

type TestAuthenticator struct {
	AuthFn   func(apiEndpoint, apiToken string) (api.KeptnInterface, error)
	ReauthFn func() (api.KeptnInterface, error)
}

func (t TestAuthenticator) Auth(apiEndpoint, apiToken string) (api.KeptnInterface, error) {
	if t.AuthFn != nil {
		return t.AuthFn(apiEndpoint, apiToken)
	}
	panic("AuthFn mock of TestAuthenticator called but not set")
}

func (t TestAuthenticator) ReAuth() (api.KeptnInterface, error) {
	if t.ReauthFn != nil {
		return t.ReauthFn()
	}
	panic("ReauthFn mock of TestAuthenticator called but not set")
}
