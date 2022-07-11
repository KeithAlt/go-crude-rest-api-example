package service

import (
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api/auth"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Products defines our service & methods
type Products struct {
	Repo *repository.Client
}

// Create creates a new product
func (svc *Products) Create(ctx *gin.Context) {
	err := auth.ValidateHeaders(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, "invalid headers", internal.ErrorInvalidArgument)
		return
	}

	jsonBytes, err := util.SerializeJSONPayload(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, "failed to serialize JSON", internal.ErrorInvalidArgument)
		return
	}

	var jsonCollection models.ModelJSONCollection
	err = json.Unmarshal(*jsonBytes, &jsonCollection.Repo)
	if err != nil {
		internal.ErrorResponse(ctx, "failed to unmarshal JSON", internal.ErrorInvalidArgument)
		return
	}

	productCollection, err := svc.Repo.Create(ctx, jsonCollection.ToModel())
	if err != nil {
		internal.ErrorResponse(ctx, "service failed to create new product", internal.ErrorServerFault)
		return
	}

	// If the client only sent one create request payload; we return one
	if len(productCollection.Repo) == 1 {
		defer ctx.JSON(http.StatusCreated, productCollection.ToJSON().Repo[0])
		return
	}

	defer ctx.JSON(http.StatusCreated, productCollection.ToJSON().Repo)
}

// Update updates a product
func (svc *Products) Update(ctx *gin.Context) {
	err := auth.ValidateHeaders(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, "invalid headers", internal.ErrorInvalidArgument)
		return
	}

	guid, err := uuid.Parse(ctx.Param("guid"))
	if err != nil {
		internal.ErrorResponse(ctx, "invalid product id", internal.ErrorInvalidArgument)
		return
	}

	var productJSON models.ProductJSON
	err = ctx.ShouldBindJSON(&productJSON)
	if err != nil {
		internal.ErrorResponse(ctx, "failed to bind JSON", internal.ErrorInvalidArgument)
		return
	}

	product, err := svc.Repo.Update(ctx, guid.String(), productJSON.ToModel())
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	defer ctx.JSON(http.StatusOK, product.ToJSON())
}

// Find returns a product by ID
func (svc *Products) Find(ctx *gin.Context) {
	err := auth.ValidateHeaders(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, "invalid headers", internal.ErrorInvalidArgument)
		return
	}

	guid, err := uuid.Parse(ctx.Param("guid"))
	if err != nil {
		internal.ErrorResponse(ctx, "invalid product id", internal.ErrorInvalidArgument)
		return
	}

	product, err := svc.Repo.Find(ctx, guid.String())
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	defer ctx.JSON(http.StatusOK, *product.ToJSON())
}

// FindAll returns all service
func (svc *Products) FindAll(ctx *gin.Context) {
	err := auth.ValidateHeaders(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, "invalid headers", internal.ErrorInvalidArgument)
		return
	}
	products, err := svc.Repo.FindAll(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, err.Error(), internal.ErrorServerFault)
		return
	}

	defer ctx.JSON(http.StatusOK, &products.ToJSON().Repo)
}

// Delete deletes a product by ID
func (svc *Products) Delete(ctx *gin.Context) {
	err := auth.ValidateHeaders(ctx)
	if err != nil {
		internal.ErrorResponse(ctx, "invalid headers", internal.ErrorInvalidArgument)
		return
	}

	guid, err := uuid.Parse(ctx.Param("guid"))
	if err != nil {
		internal.ErrorResponse(ctx, "invalid product id", internal.ErrorInvalidArgument)
		return
	}

	product, err := svc.Repo.Delete(ctx, guid.String())
	if err != nil {
		return
	}

	defer ctx.JSON(http.StatusOK, product.ToJSON())
}
