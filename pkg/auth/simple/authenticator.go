package simple

import (
	"fmt"
	"github.com/keptn/cli2/pkg/authstore"
	api "github.com/keptn/go-utils/pkg/api/utils"
)

type DefaultAuthenticator struct {
	authStore authstore.AuthStore
}

func New(authstore authstore.AuthStore) *DefaultAuthenticator {
	return &DefaultAuthenticator{authStore: authstore}
}
func (a *DefaultAuthenticator) Auth(apiEndpoint, apiToken string) (api.KeptnInterface, error) {
	apiSet, err := api.New(apiEndpoint, api.WithAuthToken(apiToken))
	if err != nil {
		return nil, fmt.Errorf("could not create API client: %w", err)
	}
	_, authErr := apiSet.AuthV1().Authenticate()
	if err != nil {
		return nil, fmt.Errorf("could not authenticate: %s", *authErr.Message)
	}

	authInfo := authstore.AuthInfo{
		APIEndpoint: apiEndpoint,
		APIToken:    apiToken,
	}
	err = a.authStore.Save(&authInfo)
	if err != nil {
		return nil, fmt.Errorf("could not save auth info: %w", err)
	}

	return apiSet, nil
}

func (a *DefaultAuthenticator) ReAuth() (api.KeptnInterface, error) {
	authInfo, err := a.authStore.Read()
	if err != nil {
		return nil, fmt.Errorf("could not authenticate: %w", err)
	}

	apiSet, err := api.New(authInfo.APIEndpoint, api.WithAuthToken(authInfo.APIToken))
	if err != nil {
		return nil, fmt.Errorf("could not create API client: %w", err)
	}
	return apiSet, nil
}
