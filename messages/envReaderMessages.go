package messages

var MissingRequiredEnv = func(key string) string {
	return "missing required env variable " + key
}

var UnableToConvertToBool = func(key string) string {
	return "unable to convert "+key+" to bool"
}

var UnableToConvertToInt = func(key string) string {
	return "unable to convert "+key+" to int"
}