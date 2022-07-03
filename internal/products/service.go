package products

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/products/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/products/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Service defines our product service
type Service struct {
	DB *postgres.Client
}

// Create creates a new product
func (svc *Service) Create(ctx *gin.Context) {
	var jsonCollection models.ModelJSONCollection
	if err := ctx.ShouldBindJSON(&jsonCollection.Repo); err != nil {
		ctx.String(http.StatusBadRequest, "invalid request payload")
		log.Fatal(err) // TODO better error handling
		return
	}

	modelCollection, err := jsonCollection.ToModel()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "the server failed to process your product JSON payload")
		log.Fatal(err) // TODO better error handling
		return
	}
	_, err = svc.DB.Create(ctx, modelCollection.Repo...)
	if err != nil {
		ctx.String(http.StatusConflict, "the server failed to create your product")
		log.Fatal(err) // TODO better error handling
		return
	}

	ctx.Status(http.StatusCreated)
}

// Update updates a product
func (svc *Service) Update(ctx *gin.Context) {
	// FIXME implement ...
	ctx.String(http.StatusOK, "update product")
}

// Find returns a product by ID
func (svc *Service) Find(ctx *gin.Context) {
	guid := ctx.Param("guid")
	product, err := svc.DB.Find(ctx, guid)
	if err != nil {
		ctx.String(http.StatusNotFound, "failed to find a product by the provided id")
		return
	}
	ctx.JSON(http.StatusOK, product.ToJSON())
}

// FindAll returns all products
func (svc *Service) FindAll(ctx *gin.Context) {
	modelCollection, err := svc.DB.FindAll(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "an internal server error occurred with your request")
		log.Fatal(err) // TODO improve error handling
		return
	}
	jsonCollection, err := modelCollection.ToJSON()
	if err != nil {
		ctx.String(http.StatusExpectationFailed, "the server failed to find any products to return")
		log.Fatal(err) // TODO improve error handling
		return
	}

	defer ctx.JSON(http.StatusOK, jsonCollection.Repo)
}

// Delete deletes a product by ID
func (svc *Service) Delete(ctx *gin.Context) {
	// TODO implement me ...
	ctx.String(http.StatusOK, "deleted product")
}
