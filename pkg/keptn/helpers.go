package keptn

import (
	"github.com/keptn/cli2/pkg/auth"
	"github.com/keptn/cli2/pkg/context"
	"github.com/keptn/cli2/pkg/mocks"
	"testing"
)

type TestKeptn struct {
	*Keptn
	TestCmdContext            *context.TestCommandContext
	KeptnProjectInterfaceMock *mocks.ProjectsV1Interface
	TestAuthenticator         *auth.TestAuthenticator
}

func NewTestKeptn(t *testing.T) *TestKeptn {
	keptnProjectInterfaceMock := &mocks.ProjectsV1Interface{}
	testAuthenticator := &auth.TestAuthenticator{}
	testContext := context.NewTestContext(t)
	k := &Keptn{
		ProjectInterface: keptnProjectInterfaceMock,
		Authenticator:    testAuthenticator,
		Context:          testContext.Context,
	}

	return &TestKeptn{
		Keptn:                     k,
		TestCmdContext:            testContext,
		KeptnProjectInterfaceMock: keptnProjectInterfaceMock,
		TestAuthenticator:         testAuthenticator,
	}
}
