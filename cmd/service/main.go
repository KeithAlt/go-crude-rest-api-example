package main

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/products/repo/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.Set()
	db := *startDatabase()
	defer startService(db)
}

// startDatabase establishes our repo connection & migrations
func startDatabase() *postgres.Client {
	cl, err := postgres.NewClient(config.DatabaseConfig)
	if err != nil {
		log.Fatal("failed to connect to repo: %w", err) // TODO improve error handling
	}

	// run our harmless migrations
	defer func(client *postgres.Client) {
		err := client.CreateTables()
		if err != nil {
			log.Fatal(err) // TODO improve error handling
		}
	}(cl)
	return cl
}

// startService begins serving our resources
func startService(client postgres.Client) {
	svc, err := api.New(&client)
	if err != nil {
		log.Fatal(err) // TODO improve error handling
	}
	r := gin.Default()
	r.GET("/products", svc.FindAll)
	r.GET("/product/:guid", svc.Find)
	r.POST("/product", svc.Create)
	r.DELETE("/products/:guid", svc.Delete)
	log.Fatal(r.Run(config.Domain)) // TODO improve error handling
}
