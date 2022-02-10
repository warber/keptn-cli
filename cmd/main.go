package main

import (
	"github.com/keptn/cli2/pkg/context"
	"github.com/keptn/cli2/pkg/keptn"
	"github.com/keptn/cli2/pkg/mocks"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cmd := createRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func createRootCommand() *cobra.Command {
	k := &keptn.Keptn{
		Context:               context.New(),
		KeptnProjectInterface: &mocks.ProjectsV1Interface{},
	}
	cmd := &cobra.Command{
		Use:     "keptn",
		Short:   "I am keptn arrrrr",
		Example: "keptn get project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(createGetCommand(k))
	cmd.AddCommand(createAuthCommand(k))
	return cmd
}
