package service

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"github.com/gin-gonic/gin"
	"log"
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
	_, err = repo.Postgres.Create(ctx, modelCollection.Repo...)
	if err != nil {
		ctx.String(http.StatusConflict, "the server failed to create your product")
		log.Fatal(err) // TODO better error handling
		return
	}

	ctx.Status(http.StatusCreated)
}

// Update updates a product
func (repo *ProductRepository) Update(ctx *gin.Context) {
	// FIXME implement ...
	ctx.String(http.StatusOK, "update product")
}

// Find returns a product by ID
func (repo *ProductRepository) Find(ctx *gin.Context) {
	guid := ctx.Param("guid")
	product, err := repo.Postgres.Find(ctx, guid)
	if err != nil {
		ctx.String(http.StatusNotFound, "failed to find a product by the provided id")
		return
	}
	ctx.JSON(http.StatusOK, product.ToJSON())
}

// FindAll returns all service
func (repo *ProductRepository) FindAll(ctx *gin.Context) {
	modelCollection, err := repo.Postgres.FindAll(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "an internal server error occurred with your request")
		log.Fatal(err) // TODO improve error handling
		return
	}
	jsonCollection, err := modelCollection.ToJSON()
	if err != nil {
		ctx.String(http.StatusExpectationFailed, "the server failed to find any service to return")
		log.Fatal(err) // TODO improve error handling
		return
	}

	defer ctx.JSON(http.StatusOK, jsonCollection.Repo)
}

// Delete deletes a product by ID
func (repo *ProductRepository) Delete(ctx *gin.Context) {
	guid := ctx.Param("guid")
	err := repo.Postgres.Delete(ctx, guid)
	if err != nil {
		ctx.String(http.StatusNotFound, "failed to find a product by the provided id")
		return
	}
	ctx.Status(http.StatusOK)
}
