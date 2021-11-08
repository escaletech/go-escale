package sns

import "github.com/aws/aws-sdk-go/service/sns/snsiface"

type Client struct {
	SNSIface snsiface.SNSAPI
}
