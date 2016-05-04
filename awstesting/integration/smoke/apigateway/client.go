//Package apigateway provides gucumber integration tests support.
package apigateway

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/apigateway"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@apigateway", func() {
		World["client"] = apigateway.New(smoke.Session)
	})
}
