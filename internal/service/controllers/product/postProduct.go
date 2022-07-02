package product

import (
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models/product"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/pkg/infrasructure/database/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO add more require constraints for request
type postJSON struct {
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

// ToModel returns a model object in model form
func (m *postJSON) ToModel() *product.Model {
	return &product.Model{
		Name:        m.Name,
		Price:       m.Price,
		Description: m.Description,
	}
}

// PostProduct returns all product
func PostProduct(db *postgres.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod postJSON
		if err := ctx.ShouldBindJSON(&prod); err != nil {
			ctx.String(http.StatusBadRequest, "invalid request payload")
			return
		}

		p := prod.ToModel()
		fmt.Println(p.GUID) // <- why does this become 0
		fmt.Println(p.ID)
		fmt.Println(p.Name)
		fmt.Println(p.Price)
		fmt.Println(p.Description)
		fmt.Println(p.UpdatedAt)
		fmt.Println(p.CreatedAt)

		// XXX FIXME - *prod.ToModel() adds an extra indices value of 0 that is the GUID being forced in
		_, err := db.Insert(ctx, *prod.ToModel())
		if err != nil {
			ctx.String(http.StatusConflict, "the server failed to create your product")
		}

		ctx.Status(http.StatusCreated)
	}
}
