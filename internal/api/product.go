package api

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/products"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/products/repo/postgres"
	"github.com/gin-gonic/gin"
)

// ServiceManager defines the global service operations
type ServiceManager interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Find(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// Handler defines our products & methods
type Handler struct {
	svc *products.Service
}

// New creates a new API service
func New(c *postgres.Client) (*Handler, error) {
	return &Handler{
		svc: &products.Service{DB: c},
	}, nil
}

// Create creates a new product
func (h *Handler) Create(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.svc.Create(ctx)
}

// Update updates a product
func (h *Handler) Update(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.svc.Find(ctx)
}

// Find returns a product by ID
func (h *Handler) Find(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.svc.Find(ctx)
}

// FindAll returns all products
func (h *Handler) FindAll(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.svc.FindAll(ctx)
}

// Delete deletes a product by ID
func (h *Handler) Delete(ctx *gin.Context) {
	// TODO implement pre-product service logic...
	defer h.svc.Delete(ctx)
}
