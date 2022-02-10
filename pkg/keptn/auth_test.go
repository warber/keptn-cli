package keptn

import (
	"github.com/keptn/cli2/pkg/auth"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuth(t *testing.T) {
	k := NewTestKeptn(t)
	k.TestAuthenticator.AuthenticatorFunc = func(options auth.AuthOptions) error { return nil }
	err := k.Auth(auth.AuthOptions{
		Endpoint: "",
		APIToken: "",
	})

	require.Nil(t, err)
}
