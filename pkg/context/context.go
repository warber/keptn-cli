package context

import (
	"io"
	"os"
)

type Context struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

func New() *Context {
	c := &Context{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}
	return c
}
