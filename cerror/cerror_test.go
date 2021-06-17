package cerror_test

import (
	"testing"

	"github.com/escaletech/go-escale/cerror"
	"github.com/stretchr/testify/assert"
)

type scenerio struct {
	errType          string
	descriptionTrue  string
	descriptionFalse string
	customError      func(error) error
	assertTrue       []func(error) bool
	assertFalse      []func(error) bool
}

func TestCError(t *testing.T) {
	scenerios := []scenerio{
		{
			errType:          "Validation",
			descriptionTrue:  "IsValidation Error",
			descriptionFalse: "IsDuplicatedError and IsNotFoundError",
			customError:      cerror.NewValidationError,
			assertTrue: []func(error) bool{
				cerror.IsValidationError,
			},
			assertFalse: []func(error) bool{
				cerror.IsDuplicatedError,
				cerror.IsNotFoundError,
			},
		},
		{
			errType:          "Duplicated",
			descriptionTrue:  "IsDuplicated Error",
			descriptionFalse: "IsValidationError and IsNotFoundError",
			customError:      cerror.NewDuplicatedError,
			assertTrue: []func(error) bool{
				cerror.IsDuplicatedError,
			},
			assertFalse: []func(error) bool{
				cerror.IsValidationError,
				cerror.IsNotFoundError,
			},
		},
		{
			errType:          "NotFound",
			descriptionTrue:  "IsNotFound Error",
			descriptionFalse: "IsDuplicatedError and IsValidationError",
			customError:      cerror.NewNotFoundError,
			assertTrue: []func(error) bool{
				cerror.IsNotFoundError,
			},
			assertFalse: []func(error) bool{
				cerror.IsDuplicatedError,
				cerror.IsValidationError,
			},
		},
	}

	var err = *new(error)
	for _, scenerio := range scenerios {
		t.Run("Error type is "+scenerio.errType+"Error", func(t *testing.T) {
			t.Run("returns true for "+scenerio.descriptionTrue+" and false for "+scenerio.descriptionFalse, func(t *testing.T) {
				customError := scenerio.customError(err)

				for _, truthy := range scenerio.assertTrue {
					assert.True(t, truthy(customError))
				}

				for _, falsy := range scenerio.assertFalse {
					assert.False(t, falsy(customError))
				}
			})
		})
	}
}
