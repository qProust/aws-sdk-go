//Package codedeploy provides gucumber integration tests support.
package codedeploy

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/codedeploy"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@codedeploy", func() {
		World["client"] = codedeploy.New(smoke.Session)
	})
}
