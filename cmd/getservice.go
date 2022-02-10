package main

import (
	"github.com/keptn/cli2/pkg/keptn"
	"github.com/spf13/cobra"
)

func createGetServicesCommand(k *keptn.Keptn) *cobra.Command {
	opts := &keptn.GetServicesOptions{}

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Get all Keptn services of a project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// execute command using k
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Project, "project", "p", "", "Keptn project name")
	return cmd
}
