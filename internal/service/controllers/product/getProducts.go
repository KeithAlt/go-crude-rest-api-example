package product

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProducts returns all products
func GetProducts(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		modelCollection, err := db.FindAll(ctx)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "an internal server error occurred with your request")
			panic(err) // TODO improve error handling
			return
		}
		jsonCollection, err := modelCollection.ToJSON()
		if err != nil {
			ctx.String(http.StatusExpectationFailed, "the server failed to find any products to return")
			panic(err) // TODO improve error handling
			return
		}

		defer ctx.JSON(http.StatusOK, jsonCollection.Repo)
	}
}
