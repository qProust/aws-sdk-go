//Package glacier provides gucumber integration tests support.
package glacier

import (
	"github.com/qProust/aws-sdk-go/awstesting/integration/smoke"
	"github.com/qProust/aws-sdk-go/service/glacier"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@glacier", func() {
		World["client"] = glacier.New(smoke.Session)
	})
}
