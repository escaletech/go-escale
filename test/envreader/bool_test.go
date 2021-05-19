package envreader_test

import (
	"testing"
	
	"github.com/escaletech/go-escale/envreader"
	"github.com/escaletech/go-escale/messages"
	testUtils "github.com/escaletech/go-escale/test/utils"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {

	t.Run("boolOrDefault", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T){
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				response := r.BoolOrDefault(testUtils.FakeEnvVarNameBool, false)

				assert.Equal(t, true, response)
				assert.Equal(t, false, r.HasErrors())
			})

			t.Run("Env doesn't have the required key", func(t *testing.T) {
				t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
					r := envreader.NewEnvReader(testUtils.GetFakeEnv)

					response := r.BoolOrDefault("invalidEnvVarName", false)

					assert.Equal(t, false, response)
					assert.Equal(t, false, r.HasErrors())
				})
			})

			t.Run("Env has the required key but its value isn't a boolean", func(t *testing.T) {
				t.Run("It should return false and an error message to the slice", func(t *testing.T) {
					r := envreader.NewEnvReader(testUtils.GetFakeEnv)
	
					response := r.BoolOrDefault(testUtils.FakeEnvVarNameString, false)

					assert.Equal(t, false, response)
					assert.Equal(t, true, r.HasErrors())
					assert.Equal(t, messages.UnableToConvertToBool(testUtils.FakeEnvVarNameString), r.Errs[0])
				})
			})
		})
	})
}