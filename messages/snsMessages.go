package messages

import "fmt"

var SnsTopicError = func(topicName string, err error) string {
	return fmt.Sprintf("Error fetching SNS topic [%s]: %s", topicName, err)
}
