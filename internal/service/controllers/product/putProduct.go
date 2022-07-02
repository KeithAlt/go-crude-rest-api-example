package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PutProduct creates a new product
func PutProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "PutProduct")
	}
}
