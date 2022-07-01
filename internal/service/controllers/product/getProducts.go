package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProducts returns all products
func GetProducts(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db.FindAll(ctx) // DEBUG
		ctx.String(http.StatusOK, "GetProducts")
	}
}
