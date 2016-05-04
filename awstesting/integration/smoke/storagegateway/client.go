//Package storagegateway provides gucumber integration tests support.
package storagegateway

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/storagegateway"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@storagegateway", func() {
		World["client"] = storagegateway.New(smoke.Session)
	})
}
