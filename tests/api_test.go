package tests

import (
	"bytes"
	"fmt"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

// TestPost tests our HTTP PUT route response & service
// XXX this function can be broken down
func TestPost(t *testing.T) {
	checkService()
	defer func() {
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

		req, err := http.NewRequest("POST", config.Host+"/products", bytes.NewBuffer(jsonData))
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
		fmt.Println(res.Body) // DEBUG
	}()
	defer killService()
}

// TestDelete tests our HTTP PUT route response & service
func TestDelete(t *testing.T) {
	checkService()
	defer func() {
		passCodes := []int{
			http.StatusOK,
			http.StatusAccepted,
			http.StatusNoContent,
		}

		target := "d4d3e181-4856-493d-8723-806400d488ea"
		url := config.Host + "/products/" + target
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
	}()
	defer killService()
}

// TestPut tests our HTTP PUT route response & service
func TestPut(t *testing.T) {
	checkService()
	defer func() {
		// TODO implement ...
	}()
	defer killService()
}

// TestGetAll tests our HTTP PUT route response & service
func TestGetAll(t *testing.T) {
	checkService()
	defer func() {
		passCodes := []int{
			http.StatusOK,
			http.StatusAccepted,
			http.StatusNoContent,
		}

		res, err := http.Get(config.Host + "/products")
		if err != nil {
			log.Println("TestGetAll failed due to an error: ", err.Error())
			t.Fail()
			return
		}

		if !checkStatusCode(res.StatusCode, passCodes) {
			t.Fail()
			log.Println("service returned an unexpected HTTP status code: ", res.Status)
			return
		}
	}()
	defer killService()
}

// TestGetByID tests our HTTP PUT route response & service
func TestGetByID(t *testing.T) {
	checkService()
	defer func() {
		// TODO: GET config.Host/product/ab595410-2f04-4b7b-a290-29e0a3fa685d
	}()
	defer killService()
}

// TestAPI tests all our service routes & responses
func TestAPI(t *testing.T) {
	checkService()
	defer func() {
		// TODO: test it all ...
	}()
	defer killService()
}
