package sns

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/escaletech/go-escale/messages"

	"github.com/pkg/errors"
)

// Create a SNS Client and test its connection
func New(awsSession *session.Session, topicsArn []string) (*Client, error) {
	c := Client{sns.New(awsSession)}

	if err := c.testConnection(topicsArn); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c Client) testConnection(topicsArn []string) error {
	res := make([]string, 0)

	for _, topicArn := range topicsArn {
		if _, err := c.GetAttributes(topicArn); err != nil {
			res = append(res, messages.SnsTopicError(topicArn, err))
		}
	}

	if len(res) > 0 {
		return errors.New(strings.Join(res, "; "))
	}

	return nil
}

// Get topic attributes
func (c Client) GetAttributes(topicArn string) (*sns.GetTopicAttributesOutput, error) {
	return c.SNSIface.GetTopicAttributes(&sns.GetTopicAttributesInput{
		TopicArn: aws.String(topicArn),
	})
}

// Publish message on topic
func (c Client) Publish(topicArn string, message string) (*sns.PublishOutput, error) {
	receiveParams := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topicArn),
	}

	return c.SNSIface.Publish(receiveParams)
}
