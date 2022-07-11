package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// TODO replace with actual headers; config.go...?
var requiredHeaders = map[string]string{
	"Secret": "EkFhJqLdPL7dCA4A", // TODO config.Secret
}

// ValidateHeaders ensures the request has the required headers
func ValidateHeaders(ctx *gin.Context) error {
	for h, v := range requiredHeaders {
		if ctx.GetHeader(h) != v {
			return fmt.Errorf("missing required header: %s", h)
		}
	}
	return nil
}
