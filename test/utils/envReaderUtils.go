package test_utils

var FakeEnvVarNameString = "fakeEnvString"
var FakeEnvVarValueString = "fakeEnvValue"
var FakeEnvVarNameBool = "fakeEnvBool"
var FakeEnvVarValueBool = "true"

func GetFakeEnv(key string) string {
	switch {
		case key == FakeEnvVarNameString:
			return FakeEnvVarValueString
		case key == FakeEnvVarNameBool :
			return FakeEnvVarValueBool
		default:
			return ""
	}
}