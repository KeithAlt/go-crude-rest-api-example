package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PutProduct creates a new product
func PutProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PutProduct")
}
