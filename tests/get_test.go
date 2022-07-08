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
	TestGetResponse(t)
	TestGetStatusCodes(t)
}

// TestGetInParallel will run all of our tests in parallel
func TestGetInParallel(t *testing.T) {
	t.Run("Test Get Code (Routine)", TestGetResponse)
	t.Run("Test Get Response (Routine)", TestGetStatusCodes)
	t.Run("Test Get All Code (Routine)", TestGetAllResponse)
	t.Run("Test Get All Response (Routine)", TestGetAllResponse)
}

// TestGetAllStatusCodes tests to ensure the response code is what we expect it to be
func TestGetAllStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		res, err := sendGetRequest("")
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

// TestGetAllResponse tests to ensure the response payload is what we expect it to be
func TestGetAllResponse(t *testing.T) {
	checkService()
	defer func() {
		res, err := sendGetRequest("")
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

// TestGetStatusCodes tests a get by id status code response
func TestGetStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		// TODO implement ...
	}()
	defer killService()
}

// TestGetResponse tests a get by id response
func TestGetResponse(t *testing.T) {
	checkService()
	defer func() {
		res, err := createTestProduct()
		if err != nil {
			t.Logf("get test failed to create a test product:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, res, err.Error())
			t.Fail()
			return
		}

		res, err = sendGetRequest("")
		if err != nil {
			t.Log("get all test HTTP handshake failed: ", err)
			t.Fail()
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Logf("get test failed to read the response payload body:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		defer res.Body.Close()

		prod := models.ProductJSON{}
		err = json.Unmarshal(body, &prod)
		if err != nil {
			t.Logf("get test failed to unmarshal the expected payload:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, string(body), err.Error())
			t.Fail()
			return
		}

		// the GUID field should never be 0
		if len(prod.GUID) <= 0 && res.StatusCode != http.StatusNoContent {
			t.Logf("get test failed to return the expected payload or status code:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}
	}()
	defer killService()
}

// sendGetRequest sends a get request
func sendGetRequest(arg string) (*http.Response, error) {
	if arg != "" {
		arg = "/" + arg
	}
	res, err := http.Get(config.Host + "/products" + arg)
	if err != nil {
		return nil, err
	}
	return res, nil
}
