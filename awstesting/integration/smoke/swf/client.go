//Package swf provides gucumber integration tests support.
package swf

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/swf"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@swf", func() {
		World["client"] = swf.New(smoke.Session)
	})
}
