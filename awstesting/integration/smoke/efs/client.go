//Package efs provides gucumber integration tests support.
package efs

import (
	"github.com/qProust/aws-sdk-go/aws"
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
