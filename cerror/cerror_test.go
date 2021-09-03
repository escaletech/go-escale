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
			descriptionFalse: "IsDuplicatedError, IsNotFoundError, UnauthorizedError and ForbiddenError",
			customError:      cerror.NewValidationError,
			assertTrue: []func(error) bool{
				cerror.IsValidationError,
			},
			assertFalse: []func(error) bool{
				cerror.IsDuplicatedError,
				cerror.IsNotFoundError,
				cerror.IsUnauthorizedError,
				cerror.IsForbiddenError,
			},
		},
		{
			errType:          "Duplicated",
			descriptionTrue:  "IsDuplicated Error",
			descriptionFalse: "IsValidationError, IsNotFoundError, UnauthorizedError and ForbiddenError",
			customError:      cerror.NewDuplicatedError,
			assertTrue: []func(error) bool{
				cerror.IsDuplicatedError,
			},
			assertFalse: []func(error) bool{
				cerror.IsValidationError,
				cerror.IsNotFoundError,
				cerror.IsUnauthorizedError,
				cerror.IsForbiddenError,
			},
		},
		{
			errType:          "NotFound",
			descriptionTrue:  "IsNotFound Error",
			descriptionFalse: "IsDuplicatedError, IsValidationError, UnauthorizedError and ForbiddenError",
			customError:      cerror.NewNotFoundError,
			assertTrue: []func(error) bool{
				cerror.IsNotFoundError,
			},
			assertFalse: []func(error) bool{
				cerror.IsDuplicatedError,
				cerror.IsValidationError,
				cerror.IsUnauthorizedError,
				cerror.IsForbiddenError,
			},
		},
		{
			errType:          "Unauthorized",
			descriptionTrue:  "Unauthorized Error",
			descriptionFalse: "IsDuplicatedError, IsValidationError, IsNotFoundError and ForbiddenError",
			customError:      cerror.NewUnauthorizedError,
			assertTrue: []func(error) bool{
				cerror.IsUnauthorizedError,
			},
			assertFalse: []func(error) bool{
				cerror.IsDuplicatedError,
				cerror.IsValidationError,
				cerror.IsNotFoundError,
				cerror.IsForbiddenError,
			},
		},
		{
			errType:          "Forbidden",
			descriptionTrue:  "IsForbidden Error",
			descriptionFalse: "IsDuplicatedError, IsValidationError, IsNotFoundError and UnauthorizedError",
			customError:      cerror.NewForbiddenError,
			assertTrue: []func(error) bool{
				cerror.IsForbiddenError,
			},
			assertFalse: []func(error) bool{
				cerror.IsDuplicatedError,
				cerror.IsValidationError,
				cerror.IsNotFoundError,
				cerror.IsUnauthorizedError,
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
