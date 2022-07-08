package tests

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"log"
	"net/http"
	"testing"
)

// expected response codes
var expectedDeleteCodes = []int{
	http.StatusOK,
	http.StatusAccepted,
	http.StatusNoContent,
}

// TestDelete will run all of our tests for the route
func TestDelete(t *testing.T) {
	TestDeleteStatusCodes(t)
	TestDeleteResponse(t)
}

// TestDeleteStatusCodes tests to ensure the response code is what we expect it to be
func TestDeleteStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		target := "d4d3e181-4856-493d-8723-806400d488ea" // TODO generate & return an actual test product
		url := config.Host + "/products/" + target
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			log.Println("delete test failed to compose HTTP post request: ", err)
			t.Fail()
			return
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		if !checkStatusCode(res.StatusCode, expectedDeleteCodes) {
			t.Fail()
			log.Println("delete test returned an unexpected HTTP status code: ", res.Status)
			return
		}
	}()
	defer killService()
}

// TestGetAllResponse tests to ensure the response payload is what we expect it to be
func TestDeleteResponse(t *testing.T) {
	checkService()
	defer func() {
		// TODO implement ...
	}()

	defer killService()
}
