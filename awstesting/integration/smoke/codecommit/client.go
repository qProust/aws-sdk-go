//Package codecommit provides gucumber integration tests support.
package codecommit

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/codecommit"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@codecommit", func() {
		World["client"] = codecommit.New(smoke.Session)
	})
}
