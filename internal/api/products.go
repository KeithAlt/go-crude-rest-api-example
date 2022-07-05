package api

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler defines our service & methods
type Handler struct {
	Repo *service.ProductRepository
}

// Create creates a new product
func (h *Handler) Create(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.Repo.Create(ctx)
}

// Update updates a product
func (h *Handler) Update(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.Repo.Update(ctx)
}

// Find returns a product by ID
func (h *Handler) Find(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.Repo.Find(ctx)
}

// FindAll returns all service
func (h *Handler) FindAll(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.Repo.FindAll(ctx)
}

// Delete deletes a product by ID
func (h *Handler) Delete(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.Repo.Delete(ctx)
}
