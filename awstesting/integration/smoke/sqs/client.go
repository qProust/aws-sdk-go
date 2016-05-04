//Package sqs provides gucumber integration tests support.
package sqs

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/sqs"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@sqs", func() {
		World["client"] = sqs.New(smoke.Session)
	})
}
