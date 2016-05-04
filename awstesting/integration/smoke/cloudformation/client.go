//Package cloudformation provides gucumber integration tests support.
package cloudformation

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/cloudformation"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudformation", func() {
		World["client"] = cloudformation.New(smoke.Session)
	})
}
