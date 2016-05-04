//Package cognitosync provides gucumber integration tests support.
package cognitosync

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/cognitosync"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cognitosync", func() {
		World["client"] = cognitosync.New(smoke.Session)
	})
}
