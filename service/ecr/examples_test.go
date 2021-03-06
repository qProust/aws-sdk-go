// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ecr_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/qProust/aws-sdk-go/aws"
	"github.com/qProust/aws-sdk-go/aws/session"
	"github.com/qProust/aws-sdk-go/service/ecr"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleECR_BatchCheckLayerAvailability() {
	svc := ecr.New(session.New())

	params := &ecr.BatchCheckLayerAvailabilityInput{
		LayerDigests: []*string{ // Required
			aws.String("BatchedOperationLayerDigest"), // Required
			// More values...
		},
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.BatchCheckLayerAvailability(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_BatchDeleteImage() {
	svc := ecr.New(session.New())

	params := &ecr.BatchDeleteImageInput{
		ImageIds: []*ecr.ImageIdentifier{ // Required
			&ecr.ImageIdentifier{ // Required
				ImageDigest: aws.String("ImageDigest"),
				ImageTag:    aws.String("ImageTag"),
			},
			// More values...
		},
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.BatchDeleteImage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_BatchGetImage() {
	svc := ecr.New(session.New())

	params := &ecr.BatchGetImageInput{
		ImageIds: []*ecr.ImageIdentifier{ // Required
			&ecr.ImageIdentifier{ // Required
				ImageDigest: aws.String("ImageDigest"),
				ImageTag:    aws.String("ImageTag"),
			},
			// More values...
		},
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.BatchGetImage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_CompleteLayerUpload() {
	svc := ecr.New(session.New())

	params := &ecr.CompleteLayerUploadInput{
		LayerDigests: []*string{ // Required
			aws.String("LayerDigest"), // Required
			// More values...
		},
		RepositoryName: aws.String("RepositoryName"), // Required
		UploadId:       aws.String("UploadId"),       // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.CompleteLayerUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_CreateRepository() {
	svc := ecr.New(session.New())

	params := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
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

func ExampleECR_DeleteRepository() {
	svc := ecr.New(session.New())

	params := &ecr.DeleteRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		Force:          aws.Bool(true),
		RegistryId:     aws.String("RegistryId"),
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

func ExampleECR_DeleteRepositoryPolicy() {
	svc := ecr.New(session.New())

	params := &ecr.DeleteRepositoryPolicyInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.DeleteRepositoryPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_DescribeRepositories() {
	svc := ecr.New(session.New())

	params := &ecr.DescribeRepositoriesInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
		RegistryId: aws.String("RegistryId"),
		RepositoryNames: []*string{
			aws.String("RepositoryName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_GetAuthorizationToken() {
	svc := ecr.New(session.New())

	params := &ecr.GetAuthorizationTokenInput{
		RegistryIds: []*string{
			aws.String("RegistryId"), // Required
			// More values...
		},
	}
	resp, err := svc.GetAuthorizationToken(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_GetDownloadUrlForLayer() {
	svc := ecr.New(session.New())

	params := &ecr.GetDownloadUrlForLayerInput{
		LayerDigest:    aws.String("LayerDigest"),    // Required
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.GetDownloadUrlForLayer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_GetRepositoryPolicy() {
	svc := ecr.New(session.New())

	params := &ecr.GetRepositoryPolicyInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.GetRepositoryPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_InitiateLayerUpload() {
	svc := ecr.New(session.New())

	params := &ecr.InitiateLayerUploadInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.InitiateLayerUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_ListImages() {
	svc := ecr.New(session.New())

	params := &ecr.ListImagesInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		MaxResults:     aws.Int64(1),
		NextToken:      aws.String("NextToken"),
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.ListImages(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_PutImage() {
	svc := ecr.New(session.New())

	params := &ecr.PutImageInput{
		ImageManifest:  aws.String("ImageManifest"),  // Required
		RepositoryName: aws.String("RepositoryName"), // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.PutImage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_SetRepositoryPolicy() {
	svc := ecr.New(session.New())

	params := &ecr.SetRepositoryPolicyInput{
		PolicyText:     aws.String("RepositoryPolicyText"), // Required
		RepositoryName: aws.String("RepositoryName"),       // Required
		Force:          aws.Bool(true),
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.SetRepositoryPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleECR_UploadLayerPart() {
	svc := ecr.New(session.New())

	params := &ecr.UploadLayerPartInput{
		LayerPartBlob:  []byte("PAYLOAD"),            // Required
		PartFirstByte:  aws.Int64(1),                 // Required
		PartLastByte:   aws.Int64(1),                 // Required
		RepositoryName: aws.String("RepositoryName"), // Required
		UploadId:       aws.String("UploadId"),       // Required
		RegistryId:     aws.String("RegistryId"),
	}
	resp, err := svc.UploadLayerPart(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
