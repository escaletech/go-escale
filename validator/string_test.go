package validator

import (
	"testing"

	"github.com/escaletech/go-escale/messages"
	"github.com/stretchr/testify/assert"
)

var v = Validator{
	Errs: make([]string, 0),
	Env:  GetFakeEnv,
}

func TestString(t *testing.T) {
	t.Run("env has the required var", func(t *testing.T) {
		t.Run("returns envvar value and no errors", func(t *testing.T) {
			response := v.StringRequired(TestEnvvar)

			assert.Equal(t, 0, len(v.Errs))
			assert.Equal(t, TestEnvvarValue, response)
		})
	})

	t.Run("env does not have the required var", func(t *testing.T) {
		t.Run("adds an error to validation slice", func(t *testing.T) {
			missingEnvvar := "WHATEVER"
			v.StringRequired(missingEnvvar)

			expectedErr := messages.RequiredEnvvar(missingEnvvar)
			assert.Equal(t, expectedErr, v.Errs[0])
		})
	})

	t.Run("env has the var", func(t *testing.T) {
		t.Run("returns envvar stored value", func(t *testing.T) {
			response := v.StringOrDefault(TestEnvvar, "DEFAULT")
			assert.Equal(t, TestEnvvarValue, response)
		})
	})

	t.Run("env doesn't have the var", func(t *testing.T) {
		t.Run("returns informed default value", func(t *testing.T) {
			response := v.StringOrDefault("WHATEVER", "DEFAULT")
			assert.Equal(t, "DEFAULT", response)
		})
	})
}
