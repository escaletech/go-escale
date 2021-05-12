package test_utils

var FakeEnvVarNameString = "fakeEnvString"
var FakeEnvVarValueString = "fakeEnvValue"
var FakeEnvVarNameBool = "fakeEnvBool"
var FakeEnvVarValueBool = "true"
var FakeEnvVarNameInt = "fakeEnvInt"
var FakeEnvVarValueInt = "13"

func GetFakeEnv(key string) string {
	switch {
	case key == FakeEnvVarNameString:
		return FakeEnvVarValueString
	case key == FakeEnvVarNameBool:
		return FakeEnvVarValueBool
	case key == FakeEnvVarNameInt:
		return FakeEnvVarValueInt
	default:
		return ""
	}
}
