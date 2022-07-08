package tests

import (
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"io/ioutil"
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
	TestPostBulkResponse(t)
}

// TestPostInParallel will run all of our tests in parallel
func TestPostInParallel(t *testing.T) {
	t.Run("Test Post Code (Routine)", TestGetAllResponse)
	t.Run("Test Post Response (Routine)", TestGetAllStatusCodes)
	t.Run("Test Bulk Post Response (Routine)", TestPostBulkResponse)
}

// TestPostStatusCodes tests to ensure the response code is what we expect it to be
func TestPostStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		res, err := createTestProduct()
		if err != nil {
			t.Log("test post failed to compose & send HTTP post request: ", err)
			t.Fail()
			return
		}

		body, _ := ioutil.ReadAll(res.Body)
		var mockProduct models.ProductJSON
		err = json.Unmarshal(body, &mockProduct)
		if err != nil {
			t.Log("test post failed to marshal the JSON payload: ", err)
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
		res, err := createTestProduct()
		if err != nil {
			t.Log("test post failed to send HTTP post request: ", err)
			t.Fail()
		}

		defer res.Body.Close()

		if !checkStatusCode(res.StatusCode, expectedPostCodes) {
			t.Log("test post request returned an unexpected error code: ", res.StatusCode)
			t.Fail()
			return
		}

		var mockProduct models.ProductJSON
		body, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &mockProduct)
		if err != nil {
			t.Log("test post failed to unmarshal the response payload: ", err)
			t.Fail()
			return
		}

		_, err = deleteProduct(mockProduct.GUID)
		if err != nil {
			t.Log("test post passed prior tests but failed to delete the mock product: ", err)
			t.Fail()
		}
	}()
	defer killService()
}

// TestPostBulkResponse tests a "bulk post" request response
func TestPostBulkResponse(t *testing.T) {
	checkService()
	defer func() {
		res, err := createTestBulkProducts()
		if err != nil {
			t.Log("test post failed to send HTTP post request: ", err)
			t.Fail()
		}

		defer res.Body.Close()

		if !checkStatusCode(res.StatusCode, expectedPostCodes) {
			t.Log("test post request returned an unexpected error code: ", res.StatusCode)
			t.Fail()
			return
		}

		var mockProducts models.ModelJSONCollection
		body, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &mockProducts)
		if err != nil {
			t.Log("test post failed to unmarshal the response payload: ", err)
			t.Fail()
			return
		}

		for _, p := range mockProducts.Repo {
			_, err = deleteProduct(p.GUID)
			if err != nil {
				t.Log("test post passed prior tests but failed to delete the mock product: ", err)
				t.Fail()
			}
		}
	}()
	defer killService()
}
