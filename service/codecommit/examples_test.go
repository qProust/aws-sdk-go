// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codecommit_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/qProust/aws-sdk-go/aws"
	"github.com/qProust/aws-sdk-go/aws/session"
	"github.com/qProust/aws-sdk-go/service/codecommit"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCodeCommit_BatchGetRepositories() {
	svc := codecommit.New(session.New())

	params := &codecommit.BatchGetRepositoriesInput{
		RepositoryNames: []*string{ // Required
			aws.String("RepositoryName"), // Required
			// More values...
		},
	}
	resp, err := svc.BatchGetRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_CreateBranch() {
	svc := codecommit.New(session.New())

	params := &codecommit.CreateBranchInput{
		BranchName:     aws.String("BranchName"),     // Required
		CommitId:       aws.String("CommitId"),       // Required
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.CreateBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_CreateRepository() {
	svc := codecommit.New(session.New())

	params := &codecommit.CreateRepositoryInput{
		RepositoryName:        aws.String("RepositoryName"), // Required
		RepositoryDescription: aws.String("RepositoryDescription"),
	}
	resp, err := svc.CreateRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_DeleteRepository() {
	svc := codecommit.New(session.New())

	params := &codecommit.DeleteRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.DeleteRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetBranch() {
	svc := codecommit.New(session.New())

	params := &codecommit.GetBranchInput{
		BranchName:     aws.String("BranchName"),
		RepositoryName: aws.String("RepositoryName"),
	}
	resp, err := svc.GetBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetCommit() {
	svc := codecommit.New(session.New())

	params := &codecommit.GetCommitInput{
		CommitId:       aws.String("ObjectId"),       // Required
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetCommit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetRepository() {
	svc := codecommit.New(session.New())

	params := &codecommit.GetRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetRepositoryTriggers() {
	svc := codecommit.New(session.New())

	params := &codecommit.GetRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"),
	}
	resp, err := svc.GetRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_ListBranches() {
	svc := codecommit.New(session.New())

	params := &codecommit.ListBranchesInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		NextToken:      aws.String("NextToken"),
	}
	resp, err := svc.ListBranches(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_ListRepositories() {
	svc := codecommit.New(session.New())

	params := &codecommit.ListRepositoriesInput{
		NextToken: aws.String("NextToken"),
		Order:     aws.String("OrderEnum"),
		SortBy:    aws.String("SortByEnum"),
	}
	resp, err := svc.ListRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_PutRepositoryTriggers() {
	svc := codecommit.New(session.New())

	params := &codecommit.PutRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"),
		Triggers: []*codecommit.RepositoryTrigger{
			&codecommit.RepositoryTrigger{ // Required
				Branches: []*string{
					aws.String("BranchName"), // Required
					// More values...
				},
				CustomData:     aws.String("RepositoryTriggerCustomData"),
				DestinationArn: aws.String("Arn"),
				Events: []*string{
					aws.String("RepositoryTriggerEventEnum"), // Required
					// More values...
				},
				Name: aws.String("RepositoryTriggerName"),
			},
			// More values...
		},
	}
	resp, err := svc.PutRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_TestRepositoryTriggers() {
	svc := codecommit.New(session.New())

	params := &codecommit.TestRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"),
		Triggers: []*codecommit.RepositoryTrigger{
			&codecommit.RepositoryTrigger{ // Required
				Branches: []*string{
					aws.String("BranchName"), // Required
					// More values...
				},
				CustomData:     aws.String("RepositoryTriggerCustomData"),
				DestinationArn: aws.String("Arn"),
				Events: []*string{
					aws.String("RepositoryTriggerEventEnum"), // Required
					// More values...
				},
				Name: aws.String("RepositoryTriggerName"),
			},
			// More values...
		},
	}
	resp, err := svc.TestRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateDefaultBranch() {
	svc := codecommit.New(session.New())

	params := &codecommit.UpdateDefaultBranchInput{
		DefaultBranchName: aws.String("BranchName"),     // Required
		RepositoryName:    aws.String("RepositoryName"), // Required
	}
	resp, err := svc.UpdateDefaultBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateRepositoryDescription() {
	svc := codecommit.New(session.New())

	params := &codecommit.UpdateRepositoryDescriptionInput{
		RepositoryName:        aws.String("RepositoryName"), // Required
		RepositoryDescription: aws.String("RepositoryDescription"),
	}
	resp, err := svc.UpdateRepositoryDescription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateRepositoryName() {
	svc := codecommit.New(session.New())

	params := &codecommit.UpdateRepositoryNameInput{
		NewName: aws.String("RepositoryName"), // Required
		OldName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.UpdateRepositoryName(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
