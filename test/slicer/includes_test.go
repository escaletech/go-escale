package slicer_test

import (
	"testing"

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
		t.Run("returns true (searched value is a struct)", func(t *testing.T) {
			refSlice := []operation{
				operationA, operationB,
			}

			response := slicer.Includes(operationA, refSlice)

			assert.Equal(t, true, response)
		})

		t.Run("returns true (searched value is a string)", func(t *testing.T) {
			refSlice := []string{"testando", "islaice"}

			response := slicer.Includes("islaice", refSlice)

			assert.Equal(t, true, response)
		})
	})

	t.Run("The slice doesn't have the searched value", func(t *testing.T) {
		t.Run("returns false (searched value is a struct)", func(t *testing.T) {
			refSlice := []operation{
				operationA,
			}

			response := slicer.Includes(operationB, refSlice)

			assert.Equal(t, false, response)
		})

		t.Run("returns false (searched value is a string)", func(t *testing.T) {
			refSlice := []string{"testando"}

			response := slicer.Includes("islaice", refSlice)

			assert.Equal(t, false, response)
		})
	})
}
