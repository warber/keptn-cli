package printer

import (
	"fmt"
	"io"
)

func PrintPlaintext(out io.Writer, v interface{}) error {
	fmt.Fprintf(out, "%v", v)
	return nil
}
