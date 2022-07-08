package tests

import (
	"bytes"
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

// expected response codes
var expectedPostCodes = []int{
	http.StatusOK,
	http.StatusAccepted,
	http.StatusConflict,
}

// TestGetAll will run all of our tests for the route
func TestPost(t *testing.T) {
	TestGetAllResponse(t)
	TestGetAllStatusCodes(t)
}

// TestPostStatusCodes tests to ensure the response code is what we expect it to be
func TestPostStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		mockJSON := createMockPayload()
		res, err := sendMockPostRequest(mockJSON)
		if err != nil {
			log.Println("testpost failed to compose & send HTTP post request: ", err)
			t.Fail()
			return
		}

		body, _ := ioutil.ReadAll(res.Body)
		var mockProduct models.ProductJSON
		err = json.Unmarshal(body, &mockProduct)
		if err != nil {
			log.Println("testpost failed to marshal the JSON payload: ", err)
			t.Fail()
			return
		}

		// TODO check that the returned JSON is what we expect it to be
		// ... expect: newly created product JSON
	}()
	defer killService()
}

// TestPostResponse tests to ensure the response code is what we expect it to be
func TestPostResponse(t *testing.T) {
	checkService()
	defer func() {
		mockJSON := createMockPayload()
		res, err := sendMockPostRequest(mockJSON)
		if err != nil {
			log.Println("testpost failed to send HTTP post request: ", err)
			t.Fail()
		}

		if !checkStatusCode(res.StatusCode, expectedPostCodes) {
			t.Fail()
			log.Println("testpost request returned an unexpected error code: ", res.StatusCode)
			return
		}
	}()
	defer killService()
}

// createMockPayload creates a mock payload
func createMockPayload() []byte {
	return []byte(`{
		"name":        "Test Product",
		"price":       500.00,
		"description": "A test product"
	}`)
}

// sendMockPostRequest sends a mock post request
func sendMockPostRequest(mockJSON []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", config.Host+"/products", bytes.NewBuffer(mockJSON))
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
	return res, nil
}
