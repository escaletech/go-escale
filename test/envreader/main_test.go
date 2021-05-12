package envreader_test

import (
	"testing"

	"github.com/escaletech/go-escale/envreader"
	"github.com/escaletech/go-escale/messages"
	testUtils "github.com/escaletech/go-escale/test/utils"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("HasErrors", func(t *testing.T) {
		t.Run("envReader errs slice is empty", func(t *testing.T) {
			t.Run("returns false", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				assert.Equal(t, false, r.HasErrors())
			})
		})

		t.Run("envReader errs slice has items", func(t *testing.T) {
			t.Run("returns false", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				r.StringRequired("WHATEVER")

				assert.Equal(t, true, r.HasErrors())
			})
		})
	})

	t.Run("GetErrors", func(t *testing.T) {
		t.Run("envReader errs slice has items", func(t *testing.T) {
			t.Run("returns items as a string", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				r.StringRequired("WHATEVER")
				r.StringRequired("WHATEVER2")

				expectedErrs := messages.MissingRequiredEnv("WHATEVER") + ", " + messages.MissingRequiredEnv("WHATEVER2")

				assert.Equal(t, expectedErrs, r.GetErrors())
			})
		})

		t.Run("envReader errs slice is empty", func(t *testing.T) {
			t.Run("returns an empty string", func(t *testing.T) {
				r := envreader.NewEnvReader(testUtils.GetFakeEnv)

				assert.Equal(t, "", r.GetErrors())
			})
		})
	})
}
