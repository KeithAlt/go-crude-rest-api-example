package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostProducts creates new products
func PostProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PostProducts")
}
