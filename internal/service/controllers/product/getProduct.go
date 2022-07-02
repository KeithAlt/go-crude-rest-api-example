package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProduct returns a product by ID
func GetProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		guid := ctx.Param("guid")
		product, err := db.FindById(ctx, guid)
		if err != nil {
			ctx.String(http.StatusNotFound, "failed to find a product by the provided id")
			panic(err) // TODO improve error handling
			return
		}
		ctx.JSON(http.StatusOK, product.ToJSON()) // FIXME this should not be JSON
	}
}
