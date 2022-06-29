package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProducts returns all products
func GetProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GetProducts")
}
