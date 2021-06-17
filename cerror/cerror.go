package cerror

import (
	"github.com/pkg/errors"
)

type ErrorType uint

const (
	NoType ErrorType = iota
	ValidationError
	DuplicatedError
	NotFoundError
)

type customError struct {
	errorType     ErrorType
	originalError error
}

func NewValidationError(err error) error {
	return customError{errorType: ValidationError, originalError: err}
}

func NewDuplicatedError(err error) error {
	return customError{errorType: DuplicatedError, originalError: err}
}

func NewNotFoundError(err error) error {
	return customError{errorType: NotFoundError, originalError: err}
}

func (error customError) Error() string {
	return error.originalError.Error()
}

func IsValidationError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == ValidationError
	}

	return false
}

func IsDuplicatedError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == DuplicatedError
	}

	return false
}

func IsNotFoundError(err error) bool {
	if customErr, ok := errors.Cause(err).(customError); ok {
		return customErr.errorType == NotFoundError
	}

	return false
}
