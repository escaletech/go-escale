package testutils

var FakeEnvVarNameString = "fakeEnvString"
var FakeEnvVarValueString = "fakeEnvValue"
var FakeEnvVarNameBool = "fakeEnvBool"
var FakeEnvVarValueBool = "true"
var FakeEnvVarNameInt = "fakeEnvInt"
var FakeEnvVarValueInt = "13"

func setFakeValues() map[string]string {
	return map[string]string{
		FakeEnvVarNameString: FakeEnvVarValueString,
		FakeEnvVarNameBool:   FakeEnvVarValueBool,
		FakeEnvVarNameInt:    FakeEnvVarValueInt,
	}
}

func GetFakeEnv(key string) string {
	return setFakeValues()[key]
}
