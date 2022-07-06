package tests

import (
	"bytes"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

var serviceStarted bool
var serviceUrl string

// checkService checks the state of our service
func checkService() {
	if !serviceStarted {
		go startService()
		serviceStarted = true
		time.Sleep(time.Second * 1)
	}
}

// startService starts our service
func startService() {
	config.Set()
	serviceUrl = config.Protocol + config.Domain
	client := *repository.Initialize()
	defer api.Serve(&client)
}

// checkStatusCode checks the returned status code to ensure it's what we expect it to be
func checkStatusCode(resCode int, passCodes []int) bool {
	for _, code := range passCodes {
		if resCode == code {
			return true
		}
	}
	return false
}

func killService() {
	_, err := http.Get(serviceUrl + "/kill")
	if err != nil {
		log.Fatal("An error occurred when killing our service: ", err)
		return
	}
	log.Println("Test service killed")
}

// TestPost tests our HTTP PUT route response & service
func TestPost(t *testing.T) {
	checkService()
	passCodes := []int{
		http.StatusOK,
		http.StatusAccepted,
		http.StatusCreated,
		http.StatusConflict,
	}

	jsonData := []byte(`{
		"name":        "Test Product",
		"price":       500.00,
		"description": "A test product"
	}`)

	req, err := http.NewRequest("POST", serviceUrl+"/products", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json;	charset=UTF-8")
	if err != nil {
		log.Println("failed to compose HTTP post request: ", err)
		t.Fail()
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		t.Fail()
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if !checkStatusCode(res.StatusCode, passCodes) {
		t.Fail()
		log.Println("the post request returned an unexpected error code: ", res.Status, "\n", string(body))
		return
	}

	log.Println("response Status:", res.Status)
	log.Println("response Headers:", res.Header)
	log.Println("response Body:", string(body))
	defer killService()
}

// TestDelete tests our HTTP PUT route response & service
func TestDelete(t *testing.T) {
	checkService()
	passCodes := []int{
		http.StatusOK,
		http.StatusAccepted,
		http.StatusNoContent,
	}

	target := "d4d3e181-4856-493d-8723-806400d488ea"
	url := serviceUrl + "/products/" + target
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println("failed to compose HTTP post request: ", err)
		t.Fail()
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if !checkStatusCode(res.StatusCode, passCodes) {
		t.Fail()
		log.Println("service returned an unexpected HTTP status code: ", res.Status)
		return
	}

	defer killService()
	// TODO: DELETE http://localhost:8080/products/ab595410-2f04-4b7b-a290-29e0a3fa685d
	// https://stackoverflow.com/questions/46310113/consume-a-delete-endpoint-from-golang
}

// TestPut tests our HTTP PUT route response & service
func TestPut(t *testing.T) {
	checkService()

	/**
	TODO: PUT http://localhost:8080/product/ab595410-2f04-4b7b-a290-29e0a3fa685d
	{
	    "name": "Updated Product",
	    "price": 500,
	    "description": "This is an updated product!"
	}
	*/
}

// TestGetAll tests our HTTP PUT route response & service
func TestGetAll(t *testing.T) {
	checkService()

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
	checkService()
	// TODO: GET http://localhost:8080/product/ab595410-2f04-4b7b-a290-29e0a3fa685d
}

// TestAPI tests all our service routes & responses
func TestAPI(t *testing.T) {
	checkService()

	// ... test various routes of our API
}
