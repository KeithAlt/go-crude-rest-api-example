package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProduct returns a product by ID
func GetProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "GetProduct")
	}
}
