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
	ErrorUnknown ErrorCode = iota
	ErrorNotFound
	ErrorInvalidArgument
	ErrorInvalidAuth
	ErrorServerFault
)

// WrapError wraps the error & throws up stack
func WrapError(stacktrace error, code ErrorCode, msg string, arg ...interface{}) error {
	return &Error{
		code:       code,
		stacktrace: stacktrace,
		msg:        fmt.Sprintf(msg, arg...),
	}
}

// UnwrapError unwraps our error
func (e *Error) UnwrapError() error {
	return e.stacktrace
}

// NewError creates a new error
func NewError(code ErrorCode, msg string, arg ...interface{}) error {
	return WrapError(nil, code, msg, arg...)
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
