package cerror

type ErrorType uint

const (
	NoType ErrorType = iota
	ValidationError
	DuplicatedError
	NotFoundError
	UnauthorizedError
	ForbiddenError
)

type customError struct {
	errorType     ErrorType
	originalError error
}
