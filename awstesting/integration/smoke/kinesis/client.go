//Package kinesis provides gucumber integration tests support.
package kinesis

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/kinesis"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@kinesis", func() {
		World["client"] = kinesis.New(smoke.Session)
	})
}
