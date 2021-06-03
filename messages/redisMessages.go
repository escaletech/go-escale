package messages

import "fmt"

var RedisConnectionError = func(identifier string, err error) string {
	return fmt.Sprintf("[%s] Can't connect to redis: %s", identifier, err)
}
