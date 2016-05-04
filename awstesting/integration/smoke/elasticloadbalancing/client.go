//Package elasticloadbalancing provides gucumber integration tests support.
package elasticloadbalancing

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/elb"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elasticloadbalancing", func() {
		World["client"] = elb.New(smoke.Session)
	})
}
