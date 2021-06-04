package envreader_test

import (
	"testing"

	"github.com/escaletech/go-escale/envreader"
	"github.com/escaletech/go-escale/messages"
	"github.com/escaletech/go-escale/test/testutils"

	"github.com/stretchr/testify/assert"
)

func init() {
	envreader.Env = testutils.GetFakeEnv
}

func TestBool(t *testing.T) {
	var er *envreader.EnvReader = envreader.New()

	t.Run("boolOrDefault", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				response := er.BoolOrDefault(testutils.FakeEnvVarNameBool, false)

				assert.Equal(t, true, response)
				assert.Equal(t, false, er.HasErrors())
			})

			t.Run("Env doesn't have the required key", func(t *testing.T) {
				t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
					response := er.BoolOrDefault("invalidEnvVarName", false)

					assert.Equal(t, false, response)
					assert.Equal(t, false, er.HasErrors())
				})
			})

			t.Run("Env has the required key but its value isn't a boolean", func(t *testing.T) {
				t.Run("It should return false and an error message to the slice", func(t *testing.T) {
					response := er.BoolOrDefault(testutils.FakeEnvVarNameString, false)

					assert.Equal(t, false, response)
					assert.Equal(t, true, er.HasErrors())
					assert.Equal(t, messages.UnableToConvertToBool(testutils.FakeEnvVarNameString), er.Errs[0])
				})
			})
		})
	})
}
