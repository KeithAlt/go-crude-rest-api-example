package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProduct returns a product by ID
func GetProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetProduct")
}
