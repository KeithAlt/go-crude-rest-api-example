package service

import (
	json2 "encoding/json" // FIXME
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Service defines the global service operations
type Service interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Find(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// ProductRepository defines our product service
type ProductRepository struct {
	Postgres *repository.Client
}

// Create creates a new product
func (repo *ProductRepository) Create(ctx *gin.Context) {
	json, err := util.SerializeJSONPayload(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	var jsonCollection models.ModelJSONCollection
	err = json2.Unmarshal(json, &jsonCollection.Repo)
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	modelCollection, err := jsonCollection.ToModel()
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	products, err := repo.Postgres.Create(ctx, modelCollection)
	if err != nil {
		if util.IsDuplicateKeyError(err) {
			ctx.Status(http.StatusConflict) // TODO replace with rest rendering
			return
		}
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	jsonProducts, err := products.ToJSON()
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	// There was only one product to create; therefor return one object
	if len(jsonProducts.Repo) == 1 {
		ctx.JSON(http.StatusCreated, jsonProducts.Repo[0])
		return
	}

	defer ctx.JSON(http.StatusCreated, jsonProducts.Repo)
}

// Update updates a product
func (repo *ProductRepository) Update(ctx *gin.Context) {
	guid := ctx.Param("guid")
	var newModelJSON models.ProductJSON
	if err := ctx.ShouldBindJSON(&newModelJSON); err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	res, err := repo.Postgres.Update(ctx, guid, newModelJSON.ToModel())
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
	}

	defer ctx.JSON(http.StatusOK, *res.ToJSON())
}

// Find returns a product by ID
func (repo *ProductRepository) Find(ctx *gin.Context) {
	guid := ctx.Param("guid")
	product, err := repo.Postgres.Find(ctx, guid)
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	defer ctx.JSON(http.StatusOK, product.ToJSON())
}

// FindAll returns all service
func (repo *ProductRepository) FindAll(ctx *gin.Context) {
	modelCollection, err := repo.Postgres.FindAll(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	jsonCollection, err := modelCollection.ToJSON()
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	defer ctx.JSON(http.StatusOK, jsonCollection.Repo)
}

// Delete deletes a product by ID
func (repo *ProductRepository) Delete(ctx *gin.Context) {
	guid := ctx.Param("guid")
	product, err := repo.Postgres.Delete(ctx, guid)
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	defer ctx.JSON(http.StatusOK, product.ToJSON())
}
