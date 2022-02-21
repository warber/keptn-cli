package keptn

import (
	"github.com/keptn/cli2/pkg/auth"
	"github.com/keptn/cli2/pkg/context"
	api "github.com/keptn/go-utils/pkg/api/utils"
)

type Keptn struct {
	*context.Context
	ProjectInterface api.ProjectsV1Interface
	Authenticator    auth.Authenticator
}
