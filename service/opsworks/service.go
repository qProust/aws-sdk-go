// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package opsworks

import (
	"github.com/qProust/aws-sdk-go/aws"
	"github.com/qProust/aws-sdk-go/aws/client"
	"github.com/qProust/aws-sdk-go/aws/client/metadata"
	"github.com/qProust/aws-sdk-go/aws/request"
	"github.com/qProust/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/qProust/aws-sdk-go/private/signer/v4"
)

// Welcome to the AWS OpsWorks API Reference. This guide provides descriptions,
// syntax, and usage examples about AWS OpsWorks actions and data types, including
// common parameters and error codes.
//
// AWS OpsWorks is an application management service that provides an integrated
// experience for overseeing the complete application lifecycle. For information
// about this product, go to the AWS OpsWorks (http://aws.amazon.com/opsworks/)
// details page.
//
//  SDKs and CLI
//
// The most common way to use the AWS OpsWorks API is by using the AWS Command
// Line Interface (CLI) or by using one of the AWS SDKs to implement applications
// in your preferred language. For more information, see:
//
//   AWS CLI (http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html)
//   AWS SDK for Java (http://docs.aws.amazon.com/AWSJavaSDK/latest/javadoc/com/amazonaws/services/opsworks/AWSOpsWorksClient.html)
//   AWS SDK for .NET (http://docs.aws.amazon.com/sdkfornet/latest/apidocs/html/N_Amazon_OpsWorks.htm)
//   AWS SDK for PHP 2 (http://docs.aws.amazon.com/aws-sdk-php-2/latest/class-Aws.OpsWorks.OpsWorksClient.html)
//   AWS SDK for Ruby (http://docs.aws.amazon.com/AWSRubySDK/latest/AWS/OpsWorks/Client.html)
//   AWS SDK for Node.js (http://aws.amazon.com/documentation/sdkforjavascript/)
//   AWS SDK for Python(Boto) (http://docs.pythonboto.org/en/latest/ref/opsworks.html)
//    Endpoints
//
// AWS OpsWorks supports only one endpoint, opsworks.us-east-1.amazonaws.com
// (HTTPS), so you must connect to that endpoint. You can then use the API to
// direct AWS OpsWorks to create stacks in any AWS Region.
//
//  Chef Versions
//
// When you call CreateStack, CloneStack, or UpdateStack we recommend you use
// the ConfigurationManager parameter to specify the Chef version. The recommended
// value for Linux stacks is currently 12 (the default is 11.4). Windows stacks
// use Chef 12.2. For more information, see Chef Versions (http://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook-chef11.html).
//
// You can specify Chef 12, 11.10, or 11.4 for your Linux stack. We recommend
// migrating your existing Linux stacks to Chef 12 as soon as possible.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OpsWorks struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "opsworks"

// New creates a new instance of the OpsWorks client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OpsWorks client from just a session.
//     svc := opsworks.New(mySession)
//
//     // Create a OpsWorks client with additional configuration
//     svc := opsworks.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *OpsWorks {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OpsWorks {
	svc := &OpsWorks{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-02-18",
				JSONVersion:   "1.1",
				TargetPrefix:  "OpsWorks_20130218",
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

// newRequest creates a new request for a OpsWorks operation and runs any
// custom request initialization.
func (c *OpsWorks) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
