//Package datapipeline provides gucumber integration tests support.
package datapipeline

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/datapipeline"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@datapipeline", func() {
		World["client"] = datapipeline.New(smoke.Session)
	})
}
