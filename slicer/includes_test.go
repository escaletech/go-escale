package slicer_test

import (
	"errors"
	"testing"

	"github.com/escaletech/go-escale/messages"
	"github.com/escaletech/go-escale/slicer"
	"github.com/stretchr/testify/assert"
)

type operation struct {
	name string
}

func TestIncludes(t *testing.T) {
	operationA := operation{
		name: "alguma",
	}

	operationB := operation{
		name: "nenhuma",
	}

	t.Run("The slice has the searched value", func(t *testing.T) {
		t.Run("returns true and no errors (searched value is a struct)", func(t *testing.T) {
			refSlice := []operation{
				operationA, operationB,
			}

			response, err := slicer.Includes(operationA, refSlice)

			assert.True(t, *response)
			assert.Nil(t, err)
		})

		t.Run("returns true and no errors (searched value is a string)", func(t *testing.T) {
			refSlice := []string{"testando", "islaice"}

			response, err := slicer.Includes("islaice", refSlice)

			assert.True(t, *response)
			assert.Nil(t, err)
		})

		t.Run("returns true and no errors (reference is an array)", func(t *testing.T) {
			var ref [2]string
			ref[0] = "abc"
			ref[1] = "def"

			response, err := slicer.Includes("def", ref)

			assert.True(t, *response)
			assert.Nil(t, err)
		})
	})

	t.Run("The slice doesn't have the searched value", func(t *testing.T) {
		t.Run("returns false and no errors (searched value is a struct)", func(t *testing.T) {
			refSlice := []operation{
				operationA,
			}

			response, err := slicer.Includes(operationB, refSlice)

			assert.False(t, *response)
			assert.Nil(t, err)
		})

		t.Run("returns false and no errors (searched value is a string)", func(t *testing.T) {
			refSlice := []string{"testando"}

			response, err := slicer.Includes("islaice", refSlice)

			assert.False(t, *response)
			assert.Nil(t, err)
		})
	})

	t.Run("Reference is not a slice or an array", func(t *testing.T) {
		t.Run("returns nil with an error for \"Reference must be a slice or an array\"", func(t *testing.T) {
			ref := "istringue"
			response, err := slicer.Includes("tantofaz", ref)

			assert.Nil(t, response)
			assert.Equal(t, errors.New(messages.RefNotSliceOrArray), err)
		})
	})
}
