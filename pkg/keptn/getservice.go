package keptn

type GetServicesOptions struct {
	Project string
}

func (o *GetServicesOptions) Validate(args []string) error {
	return nil
}
