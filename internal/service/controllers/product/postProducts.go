package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostProducts creates new products
func PostProducts(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "PostProducts")
	}
}
