package envreader_test

import (
	"testing"

	"github.com/escaletech/go-escale/envreader"
	"github.com/escaletech/go-escale/messages"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("HasErrors", func(t *testing.T) {
		t.Run("envReader errs slice is empty", func(t *testing.T) {
			t.Run("returns false", func(t *testing.T) {
				er := envreader.New()

				assert.Equal(t, false, er.HasErrors())
			})
		})

		t.Run("envReader errs slice has items", func(t *testing.T) {
			t.Run("returns false", func(t *testing.T) {
				er := envreader.New()

				er.StringRequired("WHATEVER")

				assert.Equal(t, true, er.HasErrors())
			})
		})
	})

	t.Run("GetErrors", func(t *testing.T) {
		t.Run("envReader errs slice has items", func(t *testing.T) {
			t.Run("returns items as a string", func(t *testing.T) {
				er := envreader.New()

				er.StringRequired("WHATEVER")
				er.StringRequired("WHATEVER2")

				expectedErrs := messages.MissingRequiredEnv("WHATEVER") + ", " + messages.MissingRequiredEnv("WHATEVER2")

				assert.Equal(t, expectedErrs, er.GetErrors())
			})
		})

		t.Run("envReader errs slice is empty", func(t *testing.T) {
			t.Run("returns an empty string", func(t *testing.T) {
				er := envreader.New()

				assert.Equal(t, "", er.GetErrors())
			})
		})
	})
}
