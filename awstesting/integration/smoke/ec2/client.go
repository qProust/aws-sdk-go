//Package ec2 provides gucumber integration tests support.
package ec2

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/ec2"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@ec2", func() {
		World["client"] = ec2.New(smoke.Session)
	})
}
