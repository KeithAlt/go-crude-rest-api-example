package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostProduct returns all products
func PostProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PostProduct")
}
