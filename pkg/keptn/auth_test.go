package keptn

import (
	"github.com/keptn/cli2/pkg/mocks"
	api "github.com/keptn/go-utils/pkg/api/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuth(t *testing.T) {
	k := NewTestKeptn(t)

	apiSetMock := &mocks.KeptnInterface{}
	apiSetMock.EXPECT().ProjectsV1().Return(&mocks.ProjectsV1Interface{})
	k.TestAuthenticator.AuthFn = func(apiEndpoint, apiToken string) (api.KeptnInterface, error) { return apiSetMock, nil }
	err := k.Auth(AuthOptions{})

	require.Nil(t, err)
}
