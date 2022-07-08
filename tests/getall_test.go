package tests

import (
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

// expected response codes
var expectedGetAllCodes = []int{
	http.StatusOK,
	http.StatusAccepted,
	http.StatusNoContent,
}

// TestGetAll will run all of our tests for the route
func TestGetAll(t *testing.T) {
	TestGetAllResponse(t)
	TestGetAllStatusCodes(t)
}

// TestGetAllStatusCodes tests to ensure the response code is what we expect it to be
func TestGetAllStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		res, err := http.Get(config.Host + "/products")
		if err != nil {
			log.Println("getall test failed due to an error: ", err.Error())
			t.Fail()
			return
		}

		if !checkStatusCode(res.StatusCode, expectedGetAllCodes) {
			log.Println("getall test returned an unexpected HTTP status code: ", res.Status)
			t.Fail()
			return
		}
	}()
	defer killService()
}

// TestDeleteResponse tests to ensure the response payload is what we expect it to be
func TestGetAllResponse(t *testing.T) {
	checkService()
	defer func() {
		res, err := http.Get(config.Host + "/products")
		if err != nil {
			log.Println("getall test failed to append JSON payload to our expected JSON array")
			t.Fail()
			return
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		collection := models.ModelJSONCollection{}
		err = json.Unmarshal(body, &collection.Repo)
		if err != nil {
			log.Println("getall test failed to marshal the returned payload ")
			t.Fail()
			return
		}

		if len(collection.Repo) <= 0 && res.StatusCode != http.StatusNoContent {
			log.Println("getall test failed to retrieve the expected payload or response code: ", res.StatusCode)
			t.Fail()
			return
		}
	}()
	defer killService()
}
