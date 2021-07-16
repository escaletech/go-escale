package sqs

import (
	awsSQS "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type QueueInterface interface {
	SendMessage(input SendMessageInput) error
	GetMessages(input ReceiveMessageInput) (*awsSQS.ReceiveMessageOutput, error)
	DeleteMessage(handle *string) error
}

type ReceiveMessageInput struct {
	// The maximum number of messages to return. Amazon SQS never returns more messages
	// than this value (however, fewer messages might be returned). Valid values:
	// 1 to 10. Default: 1.
	MaxNumberOfMessages *int64

	// The duration (in seconds) that the received messages are hidden from subsequent
	// retrieve requests after being retrieved by a ReceiveMessage request.
	VisibilityTimeout *int64

	// The duration (in seconds) for which the call waits for a message to arrive
	// in the queue before returning. If a message is available, the call returns
	// sooner than WaitTimeSeconds. If no messages are available and the wait time
	// expires, the call returns successfully with an empty list of messages.
	//
	// To avoid HTTP errors, ensure that the HTTP response timeout for ReceiveMessage
	// requests is longer than the WaitTimeSeconds parameter. For example, with
	// the Java SDK, you can set HTTP transport settings using the NettyNioAsyncHttpClient
	// (https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/http/nio/netty/NettyNioAsyncHttpClient.html)
	// for asynchronous clients, or the ApacheHttpClient (https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/http/apache/ApacheHttpClient.html)
	WaitTimeSeconds *int64
}

type Client struct {
	Queue  sqsiface.SQSAPI
	URL    *string
	IsFIFO bool
}

type SendMessageInput struct {
	// The length of time, in seconds, for which to delay a specific message. Valid
	// values: 0 to 900. Maximum: 15 minutes. Messages with a positive DelaySeconds
	// value become available for processing after the delay period is finished.
	// If you don't specify a value, the default value for the queue applies.
	//
	// When you set FifoQueue, you can't set DelaySeconds per message. You can set
	// this parameter only on a queue level.
	DelaySeconds int64

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The tag that specifies that a message belongs to a specific message group.
	// Messages that belong to the same message group are processed in a FIFO manner
	// (however, messages in different message groups might be processed out of
	// order). To interleave multiple ordered streams within a single queue, use
	// MessageGroupId values (for example, session data for multiple users). In
	// this scenario, multiple consumers can process the queue, but the session
	// data of each user is processed in a FIFO fashion.
	//
	//    * You must associate a non-empty MessageGroupId with a message. If you
	//    don't provide a MessageGroupId, the action fails.
	//
	//    * ReceiveMessage might return messages with multiple MessageGroupId values.
	//    For each MessageGroupId, the messages are sorted by time sent. The caller
	//    can't specify a MessageGroupId.
	//
	// The length of MessageGroupId is 128 characters. Valid values: alphanumeric
	// characters and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using MessageGroupId, see Using the MessageGroupId
	// Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagegroupid-property.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// MessageGroupId is required for FIFO queues. You can't use it for Standard
	// queues.
	MessageGroupId *string

	// The message to send. The minimum size is one character. The maximum size
	// is 256 KB.
	//
	// A message can include only XML, JSON, and unformatted text. The following
	// Unicode characters are allowed:
	//
	// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
	//
	// Any characters not included in this list will be rejected. For more information,
	// see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets).
	//
	// MessageBody is a required field
	Body string
}
