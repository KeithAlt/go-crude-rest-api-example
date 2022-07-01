package product

import (
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProducts returns all products
func GetProducts(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// FIXME: Does not work, collection fails to send
		collection, err := db.FindAll(ctx)
		if err != nil {
			ctx.String(http.StatusExpectationFailed, "an internal server error occurred with your request")
			panic(err) // TODO Improve error handling
			return
		}

		fmt.Println(collection) // FIXME debug
		ctx.JSON(http.StatusOK, collection)
	}
}
