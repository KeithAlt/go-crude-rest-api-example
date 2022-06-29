package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteProduct deletes a product by ID
func DeleteProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "deleted product")
}
