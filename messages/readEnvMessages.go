package messages

var MissingRequiredEnv = func(key string) string {
	return "missing required env variable " + key
}
