package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Error defines our custom error type
type Error struct {
	stacktrace error
	msg        string
	code       ErrorCode
}

// ErrorCode defines our internal error type
type ErrorCode uint8

// defines our internal error codes
const (
	ErrorUnknown ErrorCode = iota
	ErrorNotFound
	ErrorInvalidArgument
	ErrorUnauthorized
	ErrorServerFault
	ErrorConflict
)

// httpResponses defines our HTTP error responses in the case of an internal error
var httpResponses = map[ErrorCode]http.ConnState{
	ErrorUnknown:         http.StatusInternalServerError,
	ErrorNotFound:        http.StatusNotFound,
	ErrorInvalidArgument: http.StatusBadRequest,
	ErrorUnauthorized:    http.StatusUnauthorized,
	ErrorServerFault:     http.StatusInternalServerError,
	ErrorConflict:        http.StatusConflict,
}

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

// ErrorResponse sends an error response to our client
// TODO render responses should be moved to api/rest ...
func ErrorResponse(ctx *gin.Context, msg string, internalErr ErrorCode) {
	httpErr := http.StatusInternalServerError
	if code, ok := httpResponses[internalErr]; ok {
		httpErr = int(code)
	}
	ctx.JSON(httpErr, gin.H{
		"msg":   msg,
		"error": internalErr,
	})
}

// HandleError handles our error accordingly
// TODO render responses should be moved to api/rest ...
func HandleError(ctx *gin.Context, msg string, err error) {
	// TODO implement ...
}
