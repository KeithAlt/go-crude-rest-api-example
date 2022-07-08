package tests

import (
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"io/ioutil"
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

// TestDeleteInParallel will run all of our tests in parallel
func TestDeleteInParallel(t *testing.T) {
	t.Run("Test Delete Code (Routine)", TestDeleteStatusCodes)
	t.Run("Test Delete Response (Routine)", TestDeleteResponse)
}

// TestDeleteStatusCodes tests to ensure the response code is what we expect it to be
func TestDeleteStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		res, err := createTestProduct()
		if err != nil {
			t.Log("delete test failed to compose & send a mock delete request: ", err)
			t.Fail()
			return
		}

		if !checkStatusCode(res.StatusCode, expectedDeleteCodes) {
			t.Logf("delete test returned an illegal status code:\n|_ StatusCode = %v\n|_ Response = %v\n", res.Status, res)
			t.Fail()
			return
		}

		if err != nil {
			t.Logf("delete test failed to unmarshal the expected payload:\n|_ StatusCode = %v\n|_ Response = %v\n\n|_ Error = %s", res.Status, res, err.Error())
			t.Fail()
			return
		}

		err = unmarshalAndDeleteProduct(res)
		if err != nil {
			t.Logf("delete test past but was unable to unmarshal the response payload:\n|_ StatusCode = %v\n|_ Response = %v\n\n|_ Error = %s", res.Status, res, err.Error())
		}
	}()
	defer killService()
}

// TestGetAllResponse tests to ensure the response payload is what we expect it to be
func TestDeleteResponse(t *testing.T) {
	checkService()
	defer func() {
		res, err := createTestProduct()
		if err != nil {
			t.Log("delete test failed to compose & send a mock delete request: ", err)
			t.Fail()
			return
		}

		if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted {
			t.Logf("delete test did not return the status code for a successful deletion:\n|_ StatusCode = %v\n|_ Response = %v\n", res.Status, res)
			t.Fail()
			return
		}

		err = unmarshalAndDeleteProduct(res)
		if err != nil {
			t.Logf("delete test failed to delete product by guid:\n|_ StatusCode = %v\n|_ Response = %v\n|_ Error = %s", res.Status, res, err.Error())
			t.Fail()
			return
		}
	}()

	defer killService()
}

// unmarshalAndDeleteProduct will unmarshal a response payload & delete it from our database
func unmarshalAndDeleteProduct(res *http.Response) error {
	var mockProduct models.ProductJSON
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &mockProduct)
	if err != nil {
		return err
	}

	_, err = deleteProduct(mockProduct.GUID)
	if err != nil {
		return err
	}
	return nil
}
