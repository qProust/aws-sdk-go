//Package cloudwatch provides gucumber integration tests support.
package cloudwatch

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/cloudwatch"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudwatch", func() {
		World["client"] = cloudwatch.New(smoke.Session)
	})
}
