package api

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"github.com/gin-gonic/gin"
	"log"
)

// Serve begins serving our resources
func Serve(cl *repository.Client) {
	svc := &Handler{
		Repo: &service.ProductRepository{Postgres: cl},
	}
	r := gin.Default()
	r.GET("/products", svc.FindAll)
	r.GET("/product/:guid", svc.Find)
	r.POST("/products", svc.Create)
	r.PUT("/product/:guid", svc.Update)
	r.DELETE("/products/:guid", svc.Delete)
	log.Fatal(r.Run(config.Domain)) // TODO improve error handling
}
