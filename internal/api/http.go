package api

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createHTTPServer creates an HTTP server instance
func createHTTPServer(router *gin.Engine) *http.Server {
	server := http.Server{
		Addr:           config.HTTPServerConfig.Addr,
		Handler:        router,
		ReadTimeout:    config.HTTPServerConfig.ReadTimeout,
		WriteTimeout:   config.HTTPServerConfig.WriteTimeout,
		MaxHeaderBytes: config.HTTPServerConfig.MaxHeaderBytes,
	}
	return &server
}
