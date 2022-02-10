package context

import (
	"bytes"
	"testing"
)

type TestCommandContext struct {
	*Context
	T   *testing.T
	out *bytes.Buffer
	err *bytes.Buffer
}

func NewTestContext(t *testing.T) *TestCommandContext {
	err := &bytes.Buffer{}
	out := &bytes.Buffer{}
	cmdCtx := New()
	cmdCtx.Err = err
	cmdCtx.Out = out

	return &TestCommandContext{
		Context: cmdCtx,
		T:       t,
		err:     err,
		out:     out,
	}
}

func (c *TestCommandContext) Output() string {
	return c.out.String()
}

func (c *TestCommandContext) ErrOutput() string {
	return c.err.String()
}
