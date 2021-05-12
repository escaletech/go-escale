package test_utils

var FakeEnvVarName = "fakeEnv"
var FakeEnvVarValue = "fakeEnvValue"

func GetFakeEnv(key string) string {
	if key == FakeEnvVarName {
		return FakeEnvVarValue
	}

	return ""
}
