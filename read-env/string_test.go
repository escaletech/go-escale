package readenv_test

import (
	"fmt"
	"testing"

	readenv "github.com/escaletech/go-escale/read-env"
)

var fakeEnvVarName = "fakeEnv"
var fakeEnvVarValue = "fakeEnvValue"

func getFakeEnv(key string) string {
	if key == fakeEnvVarName {
		return fakeEnvVarValue
	} 

	return ""
}

func TestString(t *testing.T) {
	var errors []string

	var r = readenv.EnvReader{
		Errs: errors,
		Env: getFakeEnv,
	}

	t.Run("StringRequired", func(t *testing.T) { 
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				response := r.StringRequired(fakeEnvVarName)
				fmt.Println(response)
			})
		} )
	})
}