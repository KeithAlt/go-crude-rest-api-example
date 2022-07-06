package tests

import (
	"bytes"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/api"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/repository"
	"log"
	"net/http"
	"time"
)

var serviceStarted bool

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
	client := *repository.Initialize()
	defer api.Serve(&client)
}

// killService shuts down our service
func killService() {
	_, err := http.Get(config.Host + "/kill")
	if err != nil {
		log.Fatal("An error occurred when killing our service: ", err)
		return
	}
	log.Println("Test service killed")
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

func createTestProduct() (*string, error) {
	jsonData := []byte(`{
		"name":        "Test Product",
		"price":       500.00,
		"description": "A test product"
	}`)

	req, err := http.NewRequest("POST", config.Host+"/products", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json;	charset=UTF-8")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return nil, nil
}
