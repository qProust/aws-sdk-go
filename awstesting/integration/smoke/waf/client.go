//Package waf provides gucumber integration tests support.
package waf

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/waf"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@waf", func() {
		World["client"] = waf.New(smoke.Session)
	})
}
