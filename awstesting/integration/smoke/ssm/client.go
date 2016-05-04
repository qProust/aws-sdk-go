//Package ssm provides gucumber integration tests support.
package ssm

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/ssm"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@ssm", func() {
		World["client"] = ssm.New(smoke.Session)
	})
}
