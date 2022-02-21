package keptn

type AuthOptions struct {
	Endpoint string
	APIToken string
}

func (o *AuthOptions) Validate() error {
	return nil
}

func (k *Keptn) Auth(opts AuthOptions) error {
	apiSet, err := k.Authenticator.Auth(opts.Endpoint, opts.APIToken)
	if err != nil {
		return err
	}
	k.ProjectInterface = apiSet.ProjectsV1()
	return nil
}
