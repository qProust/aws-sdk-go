// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package support_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/qProust/aws-sdk-go/aws"
	"github.com/qProust/aws-sdk-go/aws/session"
	"github.com/qProust/aws-sdk-go/service/support"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSupport_AddAttachmentsToSet() {
	svc := support.New(session.New())

	params := &support.AddAttachmentsToSetInput{
		Attachments: []*support.Attachment{ // Required
			{ // Required
				Data:     []byte("PAYLOAD"),
				FileName: aws.String("FileName"),
			},
			// More values...
		},
		AttachmentSetId: aws.String("AttachmentSetId"),
	}
	resp, err := svc.AddAttachmentsToSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_AddCommunicationToCase() {
	svc := support.New(session.New())

	params := &support.AddCommunicationToCaseInput{
		CommunicationBody: aws.String("CommunicationBody"), // Required
		AttachmentSetId:   aws.String("AttachmentSetId"),
		CaseId:            aws.String("CaseId"),
		CcEmailAddresses: []*string{
			aws.String("CcEmailAddress"), // Required
			// More values...
		},
	}
	resp, err := svc.AddCommunicationToCase(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_CreateCase() {
	svc := support.New(session.New())

	params := &support.CreateCaseInput{
		CommunicationBody: aws.String("CommunicationBody"), // Required
		Subject:           aws.String("Subject"),           // Required
		AttachmentSetId:   aws.String("AttachmentSetId"),
		CategoryCode:      aws.String("CategoryCode"),
		CcEmailAddresses: []*string{
			aws.String("CcEmailAddress"), // Required
			// More values...
		},
		IssueType:    aws.String("IssueType"),
		Language:     aws.String("Language"),
		ServiceCode:  aws.String("ServiceCode"),
		SeverityCode: aws.String("SeverityCode"),
	}
	resp, err := svc.CreateCase(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeAttachment() {
	svc := support.New(session.New())

	params := &support.DescribeAttachmentInput{
		AttachmentId: aws.String("AttachmentId"), // Required
	}
	resp, err := svc.DescribeAttachment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeCases() {
	svc := support.New(session.New())

	params := &support.DescribeCasesInput{
		AfterTime:  aws.String("AfterTime"),
		BeforeTime: aws.String("BeforeTime"),
		CaseIdList: []*string{
			aws.String("CaseId"), // Required
			// More values...
		},
		DisplayId:             aws.String("DisplayId"),
		IncludeCommunications: aws.Bool(true),
		IncludeResolvedCases:  aws.Bool(true),
		Language:              aws.String("Language"),
		MaxResults:            aws.Int64(1),
		NextToken:             aws.String("NextToken"),
	}
	resp, err := svc.DescribeCases(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeCommunications() {
	svc := support.New(session.New())

	params := &support.DescribeCommunicationsInput{
		CaseId:     aws.String("CaseId"), // Required
		AfterTime:  aws.String("AfterTime"),
		BeforeTime: aws.String("BeforeTime"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeCommunications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeServices() {
	svc := support.New(session.New())

	params := &support.DescribeServicesInput{
		Language: aws.String("Language"),
		ServiceCodeList: []*string{
			aws.String("ServiceCode"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeServices(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeSeverityLevels() {
	svc := support.New(session.New())

	params := &support.DescribeSeverityLevelsInput{
		Language: aws.String("Language"),
	}
	resp, err := svc.DescribeSeverityLevels(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeTrustedAdvisorCheckRefreshStatuses() {
	svc := support.New(session.New())

	params := &support.DescribeTrustedAdvisorCheckRefreshStatusesInput{
		CheckIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrustedAdvisorCheckRefreshStatuses(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeTrustedAdvisorCheckResult() {
	svc := support.New(session.New())

	params := &support.DescribeTrustedAdvisorCheckResultInput{
		CheckId:  aws.String("String"), // Required
		Language: aws.String("String"),
	}
	resp, err := svc.DescribeTrustedAdvisorCheckResult(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeTrustedAdvisorCheckSummaries() {
	svc := support.New(session.New())

	params := &support.DescribeTrustedAdvisorCheckSummariesInput{
		CheckIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrustedAdvisorCheckSummaries(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_DescribeTrustedAdvisorChecks() {
	svc := support.New(session.New())

	params := &support.DescribeTrustedAdvisorChecksInput{
		Language: aws.String("String"), // Required
	}
	resp, err := svc.DescribeTrustedAdvisorChecks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_RefreshTrustedAdvisorCheck() {
	svc := support.New(session.New())

	params := &support.RefreshTrustedAdvisorCheckInput{
		CheckId: aws.String("String"), // Required
	}
	resp, err := svc.RefreshTrustedAdvisorCheck(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSupport_ResolveCase() {
	svc := support.New(session.New())

	params := &support.ResolveCaseInput{
		CaseId: aws.String("CaseId"),
	}
	resp, err := svc.ResolveCase(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
