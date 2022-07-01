package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models/product"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// PostProduct returns all products
func PostProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product product.ProductJSON
		if err := ctx.ShouldBindJSON(&product); err != nil {
			_, err := db.Insert(ctx, product)
			if err != nil {
				log.Fatal(err) // FIXME: bad handling >:(
			}
		}
		ctx.String(http.StatusOK, "PostProduct")
	}
}
