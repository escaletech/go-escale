package messages

import "fmt"

func RequiredEnvvar(env string) string {
	return fmt.Sprint("missing required envvar %V", env)
}
