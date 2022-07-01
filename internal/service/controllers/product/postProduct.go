package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models/product"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostProduct returns all products
func PostProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod product.ModelJSON
		if err := ctx.ShouldBindJSON(&prod); err != nil {
			_, err := db.Insert(ctx, prod)
			if err != nil {
				ctx.String(http.StatusConflict, "the server failed to create your product")
			}
		}
		ctx.Status(http.StatusCreated)
	}
}
