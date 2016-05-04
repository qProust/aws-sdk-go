// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package stsiface provides an interface for the AWS Security Token Service.
package stsiface

import (
	"github.com/qProust/aws-sdk-go/aws/request"
	"github.com/qProust/aws-sdk-go/service/sts"
)

// STSAPI is the interface type for sts.STS.
type STSAPI interface {
	AssumeRoleRequest(*sts.AssumeRoleInput) (*request.Request, *sts.AssumeRoleOutput)

	AssumeRole(*sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error)

	AssumeRoleWithSAMLRequest(*sts.AssumeRoleWithSAMLInput) (*request.Request, *sts.AssumeRoleWithSAMLOutput)

	AssumeRoleWithSAML(*sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error)

	AssumeRoleWithWebIdentityRequest(*sts.AssumeRoleWithWebIdentityInput) (*request.Request, *sts.AssumeRoleWithWebIdentityOutput)

	AssumeRoleWithWebIdentity(*sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error)

	DecodeAuthorizationMessageRequest(*sts.DecodeAuthorizationMessageInput) (*request.Request, *sts.DecodeAuthorizationMessageOutput)

	DecodeAuthorizationMessage(*sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error)

	GetCallerIdentityRequest(*sts.GetCallerIdentityInput) (*request.Request, *sts.GetCallerIdentityOutput)

	GetCallerIdentity(*sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error)

	GetFederationTokenRequest(*sts.GetFederationTokenInput) (*request.Request, *sts.GetFederationTokenOutput)

	GetFederationToken(*sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error)

	GetSessionTokenRequest(*sts.GetSessionTokenInput) (*request.Request, *sts.GetSessionTokenOutput)

	GetSessionToken(*sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error)
}

var _ STSAPI = (*sts.STS)(nil)
