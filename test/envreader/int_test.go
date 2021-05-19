package envreader_test

import (
	"testing"
	
	"github.com/escaletech/go-escale/envreader"
	"github.com/escaletech/go-escale/messages"
	testUtils "github.com/escaletech/go-escale/test/utils"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {

	t.Run("intOrDefault", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T){
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				response := r.IntOrDefault(testUtils.FakeEnvVarNameInt, 21)

				assert.Equal(t, 13, response)
				assert.Equal(t, false, r.HasErrors())
			})

			t.Run("Env doesn't have the required key", func(t *testing.T) {
				t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
					r := envreader.NewEnvReader(testUtils.GetFakeEnv)

					response := r.IntOrDefault("invalidEnvVarName", 21)

					assert.Equal(t, 21, response)
					assert.Equal(t, false, r.HasErrors())
				})
			})

			t.Run("Env has the required key but its value isn't a integer", func(t *testing.T) {
				t.Run("It should return 0 and an error message to the slice", func(t *testing.T) {
					r := envreader.NewEnvReader(testUtils.GetFakeEnv)
	
					response := r.IntOrDefault(testUtils.FakeEnvVarNameString, 21)

					assert.Equal(t, 0, response)
					assert.Equal(t, true, r.HasErrors())
					assert.Equal(t, messages.UnableToConvertToInt(testUtils.FakeEnvVarNameString), r.Errs[0])
				})
			})
		})
	})
}