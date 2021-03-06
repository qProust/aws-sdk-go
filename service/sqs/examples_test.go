// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sqs_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/qProust/aws-sdk-go/aws"
	"github.com/qProust/aws-sdk-go/aws/session"
	"github.com/qProust/aws-sdk-go/service/sqs"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSQS_AddPermission() {
	svc := sqs.New(session.New())

	params := &sqs.AddPermissionInput{
		AWSAccountIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Actions: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Label:    aws.String("String"), // Required
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.AddPermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_ChangeMessageVisibility() {
	svc := sqs.New(session.New())

	params := &sqs.ChangeMessageVisibilityInput{
		QueueUrl:          aws.String("String"), // Required
		ReceiptHandle:     aws.String("String"), // Required
		VisibilityTimeout: aws.Int64(1),         // Required
	}
	resp, err := svc.ChangeMessageVisibility(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_ChangeMessageVisibilityBatch() {
	svc := sqs.New(session.New())

	params := &sqs.ChangeMessageVisibilityBatchInput{
		Entries: []*sqs.ChangeMessageVisibilityBatchRequestEntry{ // Required
			{ // Required
				Id:                aws.String("String"), // Required
				ReceiptHandle:     aws.String("String"), // Required
				VisibilityTimeout: aws.Int64(1),
			},
			// More values...
		},
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.ChangeMessageVisibilityBatch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_CreateQueue() {
	svc := sqs.New(session.New())

	params := &sqs.CreateQueueInput{
		QueueName: aws.String("String"), // Required
		Attributes: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_DeleteMessage() {
	svc := sqs.New(session.New())

	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String("String"), // Required
		ReceiptHandle: aws.String("String"), // Required
	}
	resp, err := svc.DeleteMessage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_DeleteMessageBatch() {
	svc := sqs.New(session.New())

	params := &sqs.DeleteMessageBatchInput{
		Entries: []*sqs.DeleteMessageBatchRequestEntry{ // Required
			{ // Required
				Id:            aws.String("String"), // Required
				ReceiptHandle: aws.String("String"), // Required
			},
			// More values...
		},
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.DeleteMessageBatch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_DeleteQueue() {
	svc := sqs.New(session.New())

	params := &sqs.DeleteQueueInput{
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.DeleteQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_GetQueueAttributes() {
	svc := sqs.New(session.New())

	params := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String("String"), // Required
		AttributeNames: []*string{
			aws.String("QueueAttributeName"), // Required
			// More values...
		},
	}
	resp, err := svc.GetQueueAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_GetQueueUrl() {
	svc := sqs.New(session.New())

	params := &sqs.GetQueueUrlInput{
		QueueName:              aws.String("String"), // Required
		QueueOwnerAWSAccountId: aws.String("String"),
	}
	resp, err := svc.GetQueueUrl(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_ListDeadLetterSourceQueues() {
	svc := sqs.New(session.New())

	params := &sqs.ListDeadLetterSourceQueuesInput{
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.ListDeadLetterSourceQueues(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_ListQueues() {
	svc := sqs.New(session.New())

	params := &sqs.ListQueuesInput{
		QueueNamePrefix: aws.String("String"),
	}
	resp, err := svc.ListQueues(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_PurgeQueue() {
	svc := sqs.New(session.New())

	params := &sqs.PurgeQueueInput{
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.PurgeQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_ReceiveMessage() {
	svc := sqs.New(session.New())

	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String("String"), // Required
		AttributeNames: []*string{
			aws.String("QueueAttributeName"), // Required
			// More values...
		},
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: []*string{
			aws.String("MessageAttributeName"), // Required
			// More values...
		},
		VisibilityTimeout: aws.Int64(1),
		WaitTimeSeconds:   aws.Int64(1),
	}
	resp, err := svc.ReceiveMessage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_RemovePermission() {
	svc := sqs.New(session.New())

	params := &sqs.RemovePermissionInput{
		Label:    aws.String("String"), // Required
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.RemovePermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_SendMessage() {
	svc := sqs.New(session.New())

	params := &sqs.SendMessageInput{
		MessageBody:  aws.String("String"), // Required
		QueueUrl:     aws.String("String"), // Required
		DelaySeconds: aws.Int64(1),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Key": { // Required
				DataType: aws.String("String"), // Required
				BinaryListValues: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				BinaryValue: []byte("PAYLOAD"),
				StringListValues: []*string{
					aws.String("String"), // Required
					// More values...
				},
				StringValue: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.SendMessage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_SendMessageBatch() {
	svc := sqs.New(session.New())

	params := &sqs.SendMessageBatchInput{
		Entries: []*sqs.SendMessageBatchRequestEntry{ // Required
			{ // Required
				Id:           aws.String("String"), // Required
				MessageBody:  aws.String("String"), // Required
				DelaySeconds: aws.Int64(1),
				MessageAttributes: map[string]*sqs.MessageAttributeValue{
					"Key": { // Required
						DataType: aws.String("String"), // Required
						BinaryListValues: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						BinaryValue: []byte("PAYLOAD"),
						StringListValues: []*string{
							aws.String("String"), // Required
							// More values...
						},
						StringValue: aws.String("String"),
					},
					// More values...
				},
			},
			// More values...
		},
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.SendMessageBatch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSQS_SetQueueAttributes() {
	svc := sqs.New(session.New())

	params := &sqs.SetQueueAttributesInput{
		Attributes: map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		QueueUrl: aws.String("String"), // Required
	}
	resp, err := svc.SetQueueAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
