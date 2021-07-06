package hasher_test

import (
	"testing"

	"github.com/escaletech/go-escale/hasher"
	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	t.Run("Input is 'aaa'", func(t *testing.T) {
		t.Run("returns the hashed value", func(t *testing.T) {
			hash := hasher.StringToSha256("aaa")
			assert.Equal(t, "9834876dcfb05cb167a5c24953eba58c4ac89b1adf57f28f2f9d09af107ee8f0", hash)
		})
	})

	t.Run("Input is 123", func(t *testing.T) {
		t.Run("returns the hashed value", func(t *testing.T) {
			hash := hasher.IntToSha256(123)
			assert.Equal(t, "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3", hash)
		})
	})
}
