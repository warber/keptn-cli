package keptn

import (
	"fmt"
	"github.com/keptn/cli2/pkg/printer"
	"github.com/keptn/go-utils/pkg/api/models"
	"github.com/pkg/errors"
)

type GetProjectsOptions struct {
	printer.PrintOptions
	ProjectName string
}

func (o *GetProjectsOptions) Validate(args []string) error {
	if len(args) == 1 {
		o.ProjectName = args[0]
	} else if len(args) > 1 {
		return errors.Errorf("only one positional argument are allowed but received multiple: %v", args)
	}
	return nil
}

func (k *Keptn) GetProjects(opts *GetProjectsOptions) ([]*models.Project, error) {
	projects, err := k.KeptnProjectInterface.GetAllProjects()
	if err != nil {
		return nil, err
	}

	if opts.ProjectName != "" {
		filteredProject := []*models.Project{}
		for _, p := range projects {
			if p.ProjectName == opts.ProjectName {
				filteredProject = append(filteredProject, p)
			}
		}
		projects = filteredProject
	}
	return projects, nil
}

func (k *Keptn) PrintProjects(projects []*models.Project, opts printer.PrintOpt) error {
	switch opts.GetFormat() {
	case printer.FormatJSON:
		for _, p := range projects {
			if err := printer.PrintJSON(k.Out, p); err != nil {
				return err
			}
			return nil
		}
		return nil
	case printer.FormatYAML:
		for _, p := range projects {
			if err := printer.PrintYAML(k.Out, p); err != nil {
				return err
			}
			return nil
		}
		return nil
	default:
		return fmt.Errorf("invalid format: %s", opts.GetFormat())
	}
}
