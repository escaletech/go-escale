package readenv_test

import (
	"testing"

	"github.com/escaletech/go-escale/messages"
	"github.com/escaletech/go-escale/readenv"

	"github.com/stretchr/testify/assert"
)

var fakeEnvVarName = "fakeEnv"
var fakeEnvVarValue = "fakeEnvValue"

func getFakeEnv(key string) string {
	if key == fakeEnvVarName {
		return fakeEnvVarValue
	} 

	return ""
}

func newEnvReader() readenv.EnvReader{
	return readenv.EnvReader{
		Env: getFakeEnv,
	}
}

func TestString(t *testing.T) {

	t.Run("StringRequired", func(t *testing.T) { 
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				r := newEnvReader()

				response := r.StringRequired(fakeEnvVarName)

				assert.Equal(t, fakeEnvVarValue, response)
				assert.Equal(t, 0, len(r.Errs) )
			})
		})

		t.Run("Env doesn't have the required value", func(t *testing.T) {
			t.Run("It should return an empty string and add an error message to the slice", func(t *testing.T) {
				r := newEnvReader()

				response := r.StringRequired("invalidEnvVarName")

				assert.Equal(t, "", response)
				assert.Equal(t, messages.MissingRequiredEnv("invalidEnvVarName"), r.Errs[0])
			})
		})
	})

	t.Run("StringOrDefault", func(t *testing.T) { 
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				r := newEnvReader()

				response := r.StringOrDefault(fakeEnvVarName, "defaultValue")

				assert.Equal(t, fakeEnvVarValue, response)
				assert.Equal(t, 0, len(r.Errs) )
			})
		})

		t.Run("Env doesn't have the required value", func(t *testing.T) {
			t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
				r := newEnvReader()

				response := r.StringOrDefault("invalidEnvVarName", "defaultValue")

				assert.Equal(t, "defaultValue", response)
				assert.Equal(t, 0, len(r.Errs))
			})
		})
	})
}