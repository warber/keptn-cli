package main

import (
	"github.com/keptn/cli2/pkg/keptn"
	"github.com/spf13/cobra"
)

func createGetCommand(k *keptn.Keptn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get  [event | project | projects | stage | stages | service | services]",
		Short: "Displays an event or Keptn entities such as project, stage, or service",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(createGetProjectsCommand(k))
	cmd.AddCommand(createGetServicesCommand(k))
	return cmd
}
