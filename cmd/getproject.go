package main

import (
	"github.com/keptn/cli2/pkg/keptn"
	"github.com/spf13/cobra"
)

func createGetProjectsCommand(k *keptn.Keptn) *cobra.Command {
	opts := &keptn.GetProjectsOptions{}

	cmd := &cobra.Command{
		Use:   "project",
		Short: "Get all Keptn projects",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.ParseFormat(); err != nil {
				return err
			}
			return opts.Validate(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := k.GetProjects(opts)
			if err != nil {
				return err
			}
			return k.PrintProjects(projects, opts)
		},
	}

	cmd.Flags().StringVarP(&opts.RawFormat, "output", "o", "table", "Output format, allowed values are: table, json, yaml")
	return cmd
}
