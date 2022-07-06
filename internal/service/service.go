package service

import (
	json2 "encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api"
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
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	var jsonCollection models.ModelJSONCollection
	err = json2.Unmarshal(json, &jsonCollection.Repo)
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	modelCollection, err := jsonCollection.ToModel()
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	_, err = repo.Postgres.Create(ctx, modelCollection.Repo...)
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	ctx.Status(http.StatusCreated)
}

// Update updates a product
func (repo *ProductRepository) Update(ctx *gin.Context) {
	guid := ctx.Param("guid")
	var newModelJSON models.ProductJSON
	if err := ctx.ShouldBindJSON(&newModelJSON); err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	curModel, err := repo.Postgres.Find(ctx, guid)
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	mergedIModel := util.MergeModelsIntoInterface(curModel, &newModelJSON)
	res, err := repo.Postgres.Update(ctx, guid, mergedIModel)
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
	}
	ctx.JSON(http.StatusOK, res.ToJSON())
}

// Find returns a product by ID
func (repo *ProductRepository) Find(ctx *gin.Context) {
	guid := ctx.Param("guid")
	product, err := repo.Postgres.Find(ctx, guid)
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	ctx.JSON(http.StatusOK, product.ToJSON())
}

// FindAll returns all service
func (repo *ProductRepository) FindAll(ctx *gin.Context) {
	modelCollection, err := repo.Postgres.FindAll(ctx)
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	jsonCollection, err := modelCollection.ToJSON()
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	defer ctx.JSON(http.StatusOK, jsonCollection.Repo)
}

// Delete deletes a product by ID
func (repo *ProductRepository) Delete(ctx *gin.Context) {
	guid := ctx.Param("guid")
	err := repo.Postgres.Delete(ctx, guid)
	if err != nil {
		api.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}
	ctx.Status(http.StatusOK)
}
