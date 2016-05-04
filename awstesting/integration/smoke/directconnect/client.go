//Package directconnect provides gucumber integration tests support.
package directconnect

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/directconnect"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@directconnect", func() {
		World["client"] = directconnect.New(smoke.Session)
	})
}
