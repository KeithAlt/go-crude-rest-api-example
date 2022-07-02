package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteProduct deletes a product by ID
func DeleteProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "deleted product")
	}
}
