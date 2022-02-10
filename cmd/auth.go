package main

import (
	"github.com/keptn/cli2/pkg/auth"
	"github.com/keptn/cli2/pkg/keptn"
	"github.com/spf13/cobra"
)

func createAuthCommand(k *keptn.Keptn) *cobra.Command {
	opts := auth.AuthOptions{}

	cmd := &cobra.Command{
		Use:   "auth --endpoint=https://api.keptn.MY.DOMAIN.COM --api-token=SECRET_TOKEN",
		Short: "Authenticates the Keptn CLI against a Keptn installation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Endpoint, "endpoint", "e", "", "The endpoint exposed by the Keptn installation (e.g., api.keptn.127.0.0.1.xip.io)")
	cmd.Flags().StringVarP(&opts.APIToken, "api-token", "a", "", "The API token to communicate with the Keptn installation")
	return cmd
}
