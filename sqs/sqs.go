package sqs

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/escaletech/go-escale/messages"
	"github.com/pkg/errors"
)

// Create SQS Client Service
func New(queueName string, sqc sqsiface.SQSAPI, config Config) (*Client, error) {
	if err := validate(config); err != nil {
		return nil, err
	}

	urlResult, err := sqc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return nil, errors.WithMessage(err, messages.SQSErrorFetchingQueue)
	}

	queueURL := urlResult.QueueUrl
	return &Client{sqc, queueURL, config}, nil
}

// Retrieve messages from SQS
func (s Client) GetMessages() (*sqs.ReceiveMessageOutput, error) {
	receiveParams := &sqs.ReceiveMessageInput{
		QueueUrl:            s.URL,
		MaxNumberOfMessages: aws.Int64(s.Config.MaxNumberOfMessages),
		VisibilityTimeout:   aws.Int64(s.Config.VisibilityTimeout),
		WaitTimeSeconds:     aws.Int64(s.Config.WaitTimeSeconds),
	}
	return s.Queue.ReceiveMessage(receiveParams)
}

// Send message to SQS
func (s Client) SendMessage(body string) error {
	sendParams := &sqs.SendMessageInput{
		MessageBody:  aws.String(body),
		QueueUrl:     s.URL,
		DelaySeconds: aws.Int64(s.Config.MessageVisibilityDelay),
	}

	if _, err := s.Queue.SendMessage(sendParams); err != nil {
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

func validate(config Config) error {
	var missingParams []string

	if config.MaxNumberOfMessages == 0 {
		missingParams = append(missingParams, "MaxNumberOfMessages")
	}

	if config.MessageVisibilityDelay == 0 {
		missingParams = append(missingParams, "MessageVisibilityDelay")
	}

	if config.VisibilityTimeout == 0 {
		missingParams = append(missingParams, "VisibilityTimeout")
	}

	if config.WaitTimeSeconds == 0 {
		missingParams = append(missingParams, "WaitTimeSeconds")
	}

	if len(missingParams) > 0 {
		return errors.New(messages.SQSMissingConfigParams + strings.Join(missingParams, ", "))
	}

	return nil
}
