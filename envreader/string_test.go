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

func TestString(t *testing.T) {
	t.Run("StringRequired", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				er := envreader.New()

				response := er.StringRequired(testutils.FakeEnvVarNameString)

				assert.Equal(t, testutils.FakeEnvVarValueString, response)
				assert.Equal(t, false, er.HasErrors())
			})
		})

		t.Run("Env doesn't have the required value", func(t *testing.T) {
			t.Run("It should return an empty string and add an error message to the slice", func(t *testing.T) {
				er := envreader.New()

				response := er.StringRequired("invalidEnvVarName")

				assert.Equal(t, "", response)
				assert.Equal(t, messages.MissingRequiredEnv("invalidEnvVarName"), er.Errs[0])
			})
		})
	})

	t.Run("StringOrDefault", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				er := envreader.New()

				response := er.StringOrDefault(testutils.FakeEnvVarNameString, "defaultValue")

				assert.Equal(t, testutils.FakeEnvVarValueString, response)
				assert.Equal(t, false, er.HasErrors())
			})
		})

		t.Run("Env doesn't have the required value", func(t *testing.T) {
			t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
				er := envreader.New()

				response := er.StringOrDefault("invalidEnvVarName", "defaultValue")

				assert.Equal(t, "defaultValue", response)
				assert.Equal(t, false, er.HasErrors())
			})
		})
	})
}
