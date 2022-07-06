package tests

import (
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"log"
	"net/http"
	"testing"
)

var serviceStarted bool

// startService starts our service
func startService() {
	config.Set()
	client := *repository.Initialize()
	defer api.Serve(&client)
	serviceStarted = true
}

// TestPut tests our HTTP PUT route response & service
func TestPut(t *testing.T) {
	if !serviceStarted {
		startService()
	}

	/**
	TODO: PUT http://localhost:8080/product/ab595410-2f04-4b7b-a290-29e0a3fa685d
	{
	    "name": "Updated Product",
	    "price": 500,
	    "description": "This is an updated product!"
	}
	*/
}

// TestDelete tests our HTTP PUT route response & service
func TestDelete(t *testing.T) {
	if !serviceStarted {
		startService()
	}
	// TODO: DELETE http://localhost:8080/products/ab595410-2f04-4b7b-a290-29e0a3fa685d
}

// TestPost tests our HTTP PUT route response & service
func TestPost(t *testing.T) {
	if !serviceStarted {
		startService()
	}
	/**
	TODO: POST http://localhost:8080/products
	{
	    "name": "Lone XXX Product",
	    "price": 500,
	    "description": "the big boi 2222"
	}
	*/
}

// TestGetAll tests our HTTP PUT route response & service
func TestGetAll(t *testing.T) {
	if !serviceStarted {
		startService()
	}

	res, err := http.Get("")
	if err != nil {
		log.Println("TestGetAll failed due to an error: ", err.Error())
		t.Fail()
		return
	}

	fmt.Println(res)
	// TODO: GET http://localhost:8080/products
}

// TestGetByID tests our HTTP PUT route response & service
func TestGetByID(t *testing.T) {
	if !serviceStarted {
		startService()
	}
	// TODO: GET http://localhost:8080/product/ab595410-2f04-4b7b-a290-29e0a3fa685d
}

// TestAPI tests all our service routes & responses
func TestAPI(t *testing.T) {
	if !serviceStarted {
		startService()
	}
	// ... test various routes of our API
}
