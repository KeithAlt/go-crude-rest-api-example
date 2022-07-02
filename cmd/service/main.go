package main

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/controllers/product"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/database/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

// init sets up our service
func init() {
	config.Set()
}

func main() {
	db := *startDatabase()
	defer startService(db)
}

// startDatabase establishes our database connection & migrations
func startDatabase() *postgres.Client {
	client, err := postgres.NewClient(config.DatabaseConfig)
	if err != nil {
		log.Fatal("failed to connect to database: %w", err)
	}

	// run our harmless migrations
	defer func(client *postgres.Client) {
		err := client.CreateTables()
		if err != nil {
			log.Fatal(err)
		}
	}(client)
	return client
}

// startService begins serving our resources
func startService(client postgres.Client) {
	router := gin.Default()
	router.GET("/products", product.GetProducts(&client))
	router.GET("/product/:guid", product.GetProduct(&client))
	router.POST("/product", product.PostProduct(&client))
	router.DELETE("/products/:guid", product.DeleteProduct(&client))
	router.PUT("/product/:guid", product.PutProduct(&client))

	log.Fatal(router.Run(config.Domain))
}
