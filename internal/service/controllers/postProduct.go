package controllers

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostProduct returns all products
func PostProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db.Insert(ctx) // FIXME debug
		ctx.String(http.StatusOK, "PostProduct")
	}
}
