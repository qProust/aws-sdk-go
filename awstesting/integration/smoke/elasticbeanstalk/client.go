//Package elasticbeanstalk provides gucumber integration tests support.
package elasticbeanstalk

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/elasticbeanstalk"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elasticbeanstalk", func() {
		World["client"] = elasticbeanstalk.New(smoke.Session)
	})
}
