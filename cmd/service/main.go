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
	db := *startDatabase()
	defer startService(db)
}

// startDatabase establishes our database connection & migrations
func startDatabase() *postgres.Client {
	client, err := postgres.NewClient(config.DatabaseConfig)
	if err != nil {
		log.Fatal("failed to connect to database: %w", err)
	}

	// run our migrations
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
	router.GET("/products", controllers.GetProducts(&client))
	router.GET("/products/:guid", controllers.GetProduct(&client))
	router.POST("/product", controllers.PostProduct(&client))
	router.DELETE("/products/:guid", controllers.DeleteProduct(&client))
	router.PUT("/product/:guid", controllers.PutProduct(&client))

	log.Fatal(router.Run(config.Domain))
}
