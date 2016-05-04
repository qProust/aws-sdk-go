package route53

import (
	"regexp"

	"github.com/qProust/aws-sdk-go/aws/client"
	"github.com/qProust/aws-sdk-go/aws/request"
	"github.com/qProust/aws-sdk-go/private/protocol/restxml"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(sanitizeURL)
	}

	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opChangeResourceRecordSets:
			r.Handlers.UnmarshalError.Remove(restxml.UnmarshalErrorHandler)
			r.Handlers.UnmarshalError.PushBack(unmarshalChangeResourceRecordSetsError)
		}
	}
}

var reSanitizeURL = regexp.MustCompile(`\/%2F\w+%2F`)

func sanitizeURL(r *request.Request) {
	r.HTTPRequest.URL.Opaque =
		reSanitizeURL.ReplaceAllString(r.HTTPRequest.URL.Opaque, "/")
}
