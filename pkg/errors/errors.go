package errors

type ErrorType string

const (
	NotFound     ErrorType = "NOT_FOUND"
	Duplicate    ErrorType = "DUPLICATE"
	InvalidInput ErrorType = "INVALID_INPUT"
	Internal     ErrorType = "INTERNAL"
)

type DomainError struct {
	Type    ErrorType
	Message string
	Code    int
}

func (e *DomainError) Error() string {
	return e.Message
}
