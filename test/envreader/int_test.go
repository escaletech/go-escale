package envreader_test

import (
	"testing"

	"github.com/escaletech/go-escale/envreader"
	"github.com/escaletech/go-escale/messages"
	"github.com/escaletech/go-escale/test/testutils"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	t.Run("intOrDefault", func(t *testing.T) {
		t.Run("Env has the required value", func(t *testing.T) {
			t.Run("It should return the variable value and error's slice should be empty", func(t *testing.T) {
				er := envreader.New()

				response := er.IntOrDefault(testutils.FakeEnvVarNameInt, 21)

				assert.Equal(t, 13, response)
				assert.Equal(t, false, er.HasErrors())
			})

			t.Run("Env doesn't have the required key", func(t *testing.T) {
				t.Run("It should return a default value and error's slice should be empty", func(t *testing.T) {
					er := envreader.New()

					response := er.IntOrDefault("invalidEnvVarName", 21)

					assert.Equal(t, 21, response)
					assert.Equal(t, false, er.HasErrors())
				})
			})

			t.Run("Env has the required key but its value isn't a integer", func(t *testing.T) {
				t.Run("It should return 0 and an error message to the slice", func(t *testing.T) {
					er := envreader.New()

					response := er.IntOrDefault(testutils.FakeEnvVarNameString, 21)

					assert.Equal(t, 0, response)
					assert.Equal(t, true, er.HasErrors())
					assert.Equal(t, messages.UnableToConvertToInt(testutils.FakeEnvVarNameString), er.Errs[0])
				})
			})
		})
	})
}
