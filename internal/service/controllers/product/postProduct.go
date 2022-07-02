package product

import (
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models/product"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostProduct returns all product
func PostProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var jsonCollection product.ModelJSONCollection
		if err := ctx.ShouldBindJSON(&jsonCollection.Repo); err != nil {
			ctx.String(http.StatusBadRequest, "invalid request payload")
			panic(err) // TODO better error handling
			return
		}

		modelCollection, err := jsonCollection.ToModel()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "the server failed to process your product JSON payload")
			panic(err) // TODO better error handling
			return
		}
		res, err := db.Insert(ctx, modelCollection.Repo...)
		if err != nil {
			ctx.String(http.StatusConflict, "the server failed to create your product")
			panic(err) // TODO better error handling
			return
		}

		fmt.Println(res) // FIXME debug

		ctx.Status(http.StatusCreated)
	}
}
