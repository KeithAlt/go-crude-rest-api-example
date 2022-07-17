package config

import (
	"github.com/gin-gonic/gin"
	"time"
)

// HTTPServer defines our configuration for the HTTP server
type HTTPServer struct {
	Addr           string `binding:"required"`
	Handler        *gin.Engine
	ReadTimeout    time.Duration `binding:"required"`
	WriteTimeout   time.Duration `binding:"required"`
	MaxHeaderBytes int           `binding:"required"`
}

// setHTTPServerConfig configures our HTTP server
func setHTTPServerConfig() {
	HTTPServerConfig = HTTPServer{
		Addr:           Domain,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
