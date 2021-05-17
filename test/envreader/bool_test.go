package envreader_test

import (
	"testing"
	
	"github.com/escaletech/go-escale/envreader"
	testUtils "github.com/escaletech/go-escale/test/utils"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {

	t.Run("boolOrDefault", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T){
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				response := r.BoolOrDefault(testUtils.FakeEnvVarNameBool, true)

				assert.Equal(t, true, response)
				assert.Equal(t, false, r.HasErrors())
			})

			t.Run("Env doesn't have the required value", func(t *testing.T) {
				t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
					r := envreader.NewEnvReader(testUtils.GetFakeEnv)
	
					response := r.BoolOrDefault("invalidEnvVarName", false)
	
					assert.Equal(t, false, response)
					assert.Equal(t, false, r.HasErrors())
				})
			})
		})
	})
}