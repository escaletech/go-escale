package validator

var TestEnvvar = "envvar"
var TestEnvvarValue = "value"

var GetFakeEnv = func(key string) string {
	if key == TestEnvvar {
		return TestEnvvarValue
	}

	return ""
}
