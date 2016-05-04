// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kinesis

import (
	"github.com/qProust/aws-sdk-go/aws"
	"github.com/qProust/aws-sdk-go/aws/client"
	"github.com/qProust/aws-sdk-go/aws/client/metadata"
	"github.com/qProust/aws-sdk-go/aws/request"
	"github.com/qProust/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/qProust/aws-sdk-go/private/signer/v4"
)

// Amazon Kinesis Streams is a managed service that scales elastically for real
// time processing of streaming big data.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type Kinesis struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "kinesis"

// New creates a new instance of the Kinesis client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Kinesis client from just a session.
//     svc := kinesis.New(mySession)
//
//     // Create a Kinesis client with additional configuration
//     svc := kinesis.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Kinesis {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *Kinesis {
	svc := &Kinesis{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-12-02",
				JSONVersion:   "1.1",
				TargetPrefix:  "Kinesis_20131202",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Kinesis operation and runs any
// custom request initialization.
func (c *Kinesis) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
