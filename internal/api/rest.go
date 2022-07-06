package api

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

// errorCodes defines our HTTP error responses in the case of an internal error
var errorCodes = map[internal.ErrorCode]http.ConnState{
	internal.ErrorUnknown:         http.StatusInternalServerError,
	internal.ErrorNotFound:        http.StatusNotFound,
	internal.ErrorInvalidArgument: http.StatusBadRequest,
	internal.ErrorUnauthorized:    http.StatusUnauthorized,
	internal.ErrorServerFault:     http.StatusInternalServerError,
}

// ErrorResponse sends an error response to our client
func ErrorResponse(ctx *gin.Context, msg string, internalErr internal.ErrorCode) {
	httpErr := http.StatusInternalServerError
	if code, ok := errorCodes[internalErr]; ok {
		httpErr = int(code)
	}
	ctx.JSON(httpErr, gin.H{
		"msg":   msg,
		"error": internalErr,
	})
}

// HandleError handles our error accordingly
func HandleError(ctx *gin.Context, msg string, err error) {
	// TODO implement ...
}
