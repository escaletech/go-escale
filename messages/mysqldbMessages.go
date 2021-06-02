package messages

import "fmt"

var DBConnectionError = func(identifier string, err error) string {
	return fmt.Sprintf("[%s] Can't connect to database: %s", identifier, err)
}
