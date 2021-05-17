package envreader_test

import (
	"testing"

	"github.com/escaletech/go-escale/envreader"
	"github.com/escaletech/go-escale/messages"
	testUtils "github.com/escaletech/go-escale/test/utils"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("StringRequired", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				response := r.StringRequired(testUtils.FakeEnvVarNameString)

				assert.Equal(t, testUtils.FakeEnvVarValueString, response)
				assert.Equal(t, false, r.HasErrors())
			})
		})

		t.Run("Env doesn't have the required value", func(t *testing.T) {
			t.Run("It should return an empty string and add an error message to the slice", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				response := r.StringRequired("invalidEnvVarName")

				assert.Equal(t, "", response)
				assert.Equal(t, messages.MissingRequiredEnv("invalidEnvVarName"), r.Errs[0])
			})
		})
	})

	t.Run("StringOrDefault", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				response := r.StringOrDefault(testUtils.FakeEnvVarNameString, "defaultValue")

				assert.Equal(t, testUtils.FakeEnvVarValueString, response)
				assert.Equal(t, false, r.HasErrors())
			})
		})

		t.Run("Env doesn't have the required value", func(t *testing.T) {
			t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				response := r.StringOrDefault("invalidEnvVarName", "defaultValue")

				assert.Equal(t, "defaultValue", response)
				assert.Equal(t, false, r.HasErrors())
			})
		})
	})
}
