//Package dynamodbstreams provides gucumber integration tests support.
package dynamodbstreams

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/dynamodbstreams"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@dynamodbstreams", func() {
		World["client"] = dynamodbstreams.New(smoke.Session)
	})
}
