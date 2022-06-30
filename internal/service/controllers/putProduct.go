package controllers

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PutProduct creates a new product
func PutProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "PutProduct")
	}
}
