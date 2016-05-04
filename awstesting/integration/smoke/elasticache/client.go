//Package elasticache provides gucumber integration tests support.
package elasticache

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/elasticache"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elasticache", func() {
		World["client"] = elasticache.New(smoke.Session)
	})
}
