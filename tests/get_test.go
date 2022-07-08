package tests

import (
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"io/ioutil"
	"net/http"
	"testing"
)

// expected response codes
var expectedGetCodes = []int{
	http.StatusOK,
	http.StatusAccepted,
	http.StatusNoContent,
}

// TestGet will run all of our tests for the route
func TestGet(t *testing.T) {
	TestGetAllResponse(t)
	TestGetAllStatusCodes(t)
}

// TestGetInParallel will run all of our tests in parallel
func TestGetInParallel(t *testing.T) {
	t.Run("Test Get Code (Routine)", TestGetAllResponse)
	t.Run("Test Get Response (Routine)", TestGetAllResponse)
}

// TestGetAllStatusCodes tests to ensure the response code is what we expect it to be
func TestGetAllStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		res, err := sendGetAllRequest()
		if err != nil {
			t.Log("get all test HTTP handshake failed: ", err)
			t.Fail()
			return
		}

		if !checkStatusCode(res.StatusCode, expectedGetCodes) {
			t.Log("get all test returned an illegal status code: ", res.StatusCode)
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
		res, err := sendGetAllRequest()
		if err != nil {
			t.Log("get all test HTTP handshake failed: ", err)
			t.Fail()
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Logf("get all test failed to read the response payload body:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		defer res.Body.Close()

		collection := models.ModelJSONCollection{}
		err = json.Unmarshal(body, &collection.Repo)
		if err != nil {
			t.Logf("get all test failed to unmarshal the expected payload:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		if len(collection.Repo) <= 0 && res.StatusCode != http.StatusNoContent {
			t.Logf("get all test failed to return the expected payload or status code:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}
	}()
	defer killService()
}

// sendGetAllRequest sends a get all request
func sendGetAllRequest() (*http.Response, error) {
	res, err := http.Get(config.Host + "/products")
	if err != nil {
		return nil, err
	}
	return res, nil
}
