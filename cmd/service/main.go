package main

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.Set()
	db := *startDatabase()
	defer startService(db)
}

// startDatabase establishes our repo connection & migrations
func startDatabase() *repository.Client {
	cl, err := repository.NewClient(config.DatabaseConfig)
	if err != nil {
		log.Fatal("failed to connect to repo: %w", err) // TODO improve error handling
	}

	// run our harmless migrations
	defer func(client *repository.Client) {
		err := client.CreateTables()
		if err != nil {
			log.Fatal(err) // TODO improve error handling
		}
	}(cl)
	return cl
}

// startService begins serving our resources
func startService(client repository.Client) {
	svc, err := api.New(&client)
	if err != nil {
		log.Fatal(err) // TODO improve error handling
	}
	r := gin.Default()
	r.GET("/products", svc.FindAll)
	r.GET("/products/:guid", svc.Find)
	r.POST("/products", svc.Create)
	r.DELETE("/products/:guid", svc.Delete)
	log.Fatal(r.Run(config.Domain)) // TODO improve error handling
}
