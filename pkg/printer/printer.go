package printer

import "github.com/pkg/errors"

type Format string

const (
	FormatJSON      Format = "json"
	FormatTable     Format = "table"
	FormatYAML      Format = "yaml"
	FormatPlaintext Format = "plaintext"
)

type Formats []Format

type PrintOpt interface {
	ParseFormat() error
	GetFormat() Format
}

type PrintOptions struct {
	RawFormat string
	Format
}

func (p *PrintOptions) ParseFormat() error {
	format := Format(p.RawFormat)
	switch format {
	case FormatTable, FormatJSON, FormatYAML, FormatPlaintext:
		p.Format = format
		return nil
	default:
		return errors.Errorf("invalid format: %s", p.RawFormat)
	}
}

func (p *PrintOptions) GetFormat() Format {
	return p.Format
}
