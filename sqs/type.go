package sqs

import (
	awsSQS "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type QueueInterface interface {
	SendMessage(body string, messageGroupId *string) error
	GetMessages() (*awsSQS.ReceiveMessageOutput, error)
	DeleteMessage(handle *string) error
}

type Config struct {
	MaxNumberOfMessages    int64
	VisibilityTimeout      int64
	WaitTimeSeconds        int64
	MessageVisibilityDelay int64
}

type Client struct {
	Queue  sqsiface.SQSAPI
	URL    *string
	Config Config
}
