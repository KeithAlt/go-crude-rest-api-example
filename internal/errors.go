package internal

import "fmt"

// Error defines our custom error type
type Error struct {
	stacktrace error
	msg        string
	code       ErrorCode
}

// ErrorCode defines the types of error codes
type ErrorCode uint8

const (
	ErrorStatusUnknown ErrorCode = iota
	ErrorStatusNotFound
	ErrorStatusInvalidArgument
)

// WrapError wraps the error & throws up stack
func WrapError(stacktrace error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		code:       code,
		stacktrace: stacktrace,
		msg:        fmt.Sprintf(format, a...),
	}
}

// UnwrapError unwraps our error
func (e *Error) UnwrapError() error {
	return e.stacktrace
}

// NewError creates a new error
func NewError(code ErrorCode, format string, a ...interface{}) error {
	return WrapError(nil, code, format, a...)
}

// Error returns our error message
func (e *Error) Error() string {
	if e.stacktrace != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.stacktrace)
	}
	return e.msg
}

// ErrorCode returns our error code
func (e *Error) ErrorCode() ErrorCode {
	return e.code
}
