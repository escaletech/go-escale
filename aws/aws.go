package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Start new AWS session
func New(awsURL, awsRegion string) *session.Session {
	sessConfig := &aws.Config{
		Region: aws.String(awsRegion),
	}
	if awsURL != "" {
		sessConfig.Endpoint = aws.String(awsURL)
	}
	sess := session.Must(session.NewSession(sessConfig))

	return sess
}
