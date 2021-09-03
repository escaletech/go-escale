package cerror

import (
	"github.com/pkg/errors"
)

// Create a validation error
func NewValidationError(err error) error {
	return customError{errorType: ValidationError, originalError: err}
}

// Create a Duplicated error
func NewDuplicatedError(err error) error {
	return customError{errorType: DuplicatedError, originalError: err}
}

// Create a Not Found error
func NewNotFoundError(err error) error {
	return customError{errorType: NotFoundError, originalError: err}
}

// Create a Unauthorized error
func NewUnauthorizedError(err error) error {
	return customError{errorType: UnauthorizedError, originalError: err}
}

// Create a Forbidden error
func NewForbiddenError(err error) error {
	return customError{errorType: ForbiddenError, originalError: err}
}

// Get error message from a custom error
func (err customError) Error() string {
	return err.originalError.Error()
}

// Check if given error has the "Validation Error" type
func IsValidationError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == ValidationError
	}

	return false
}

// Check if given error has the "Is Duplicated Error" type
func IsDuplicatedError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == DuplicatedError
	}

	return false
}

// Check if given error has the "Not Found Error" type
func IsNotFoundError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == NotFoundError
	}

	return false
}

// Check if given error has the "Unauthorized Error" type
func IsUnauthorizedError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == UnauthorizedError
	}

	return false
}

// Check if given error has the "Forbidden Error" type
func IsForbiddenError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == ForbiddenError
	}

	return false
}
