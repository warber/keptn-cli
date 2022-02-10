package keptn

import "github.com/keptn/cli2/pkg/auth"

func (k *Keptn) Auth(opts auth.AuthOptions) error {
	err := k.Authenticator.Auth(opts)
	if err != nil {
		return err
	}
	return nil
}
