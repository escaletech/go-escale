package sqs

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/escaletech/go-escale/messages"
	"github.com/pkg/errors"
)

// Create SQS Client
func New(queueName string, isFIFO bool, sqc sqsiface.SQSAPI) (*Client, error) {
	urlResult, err := sqc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return nil, errors.WithMessage(err, messages.SQSErrorFetchingQueue)
	}

	queueURL := urlResult.QueueUrl
	return &Client{sqc, queueURL, isFIFO}, nil
}

// Retrieve messages from SQS
func (s Client) GetMessages(input ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	if err := validateReceiveMessageInput(input); err != nil {
		return nil, err
	}

	receiveParams := &sqs.ReceiveMessageInput{
		QueueUrl:            s.URL,
		MaxNumberOfMessages: aws.Int64(*input.MaxNumberOfMessages),
		VisibilityTimeout:   aws.Int64(*input.VisibilityTimeout),
		WaitTimeSeconds:     aws.Int64(*input.WaitTimeSeconds),
	}
	return s.Queue.ReceiveMessage(receiveParams)
}

// Send message to SQS
func (s Client) SendMessage(input SendMessageInput) error {
	sendParams := sqs.SendMessageInput{
		MessageBody: aws.String(input.Body),
		QueueUrl:    s.URL,
	}

	if !s.IsFIFO {
		sendParams.DelaySeconds = aws.Int64(input.DelaySeconds)
	}

	if input.MessageGroupId != nil {
		sendParams.MessageGroupId = input.MessageGroupId
	}

	if _, err := s.Queue.SendMessage(&sendParams); err != nil {
		return err
	}

	return nil
}

// Delete message on SQS
func (s Client) DeleteMessage(handle *string) error {
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      s.URL,
		ReceiptHandle: handle,
	}

	if _, err := s.Queue.DeleteMessage(deleteParams); err != nil {
		return err
	}

	return nil
}

func validateReceiveMessageInput(input ReceiveMessageInput) error {
	var missingParams []string

	if input.MaxNumberOfMessages == nil {
		missingParams = append(missingParams, "MaxNumberOfMessages")
	}

	if input.VisibilityTimeout == nil {
		missingParams = append(missingParams, "VisibilityTimeout")
	}

	if input.WaitTimeSeconds == nil {
		missingParams = append(missingParams, "WaitTimeSeconds")
	}

	if len(missingParams) > 0 {
		return errors.New(messages.SQSMissingConfigParams + strings.Join(missingParams, ", "))
	}

	return nil
}
