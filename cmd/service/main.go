package main

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/infrasructure/database/postgres"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

// init sets up our service
func init() {
	config.Set()
}

func main() {
	startDatabase()
	defer startService()
}

// startDatabase establishes our database connection & migrations
func startDatabase() {
	client, err := postgres.NewClient(config.DatabaseConfig)
	if err != nil {
		log.Fatal("failed to connect to database: %w", err)
	}

	defer func(client *postgres.Client) {
		err := client.CreateTables()
		if err != nil {
			log.Fatal(err)
		}
	}(client)
}

// startService begins serving our resources
func startService() {
	router := gin.Default()
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:guid", controllers.GetProduct)
	router.POST("/products", controllers.PostProduct)
	router.DELETE("/products/:guid", controllers.DeleteProduct)
	router.PUT("/product/:guid", controllers.PutProduct)

	log.Fatal(router.Run(config.Domain))
}
