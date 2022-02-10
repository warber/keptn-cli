package printer

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"io"
)

func PrintYAML(out io.Writer, v interface{}) error {
	b, err := yaml.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "could not marshal value to yaml")
	}
	fmt.Fprintln(out, string(b))
	return nil
}
