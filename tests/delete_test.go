package tests

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/tests/testkit"
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
func TestAllDelete(t *testing.T) {
	TestDelete(t)
}

// TestDeleteInParallel will run all of our tests in parallel
func TestDeleteInParallel(t *testing.T) {
	t.Run("TEST DELETE Routine", TestDelete)
}

// TestDelete tests to ensure the response code is what we expect it to be
func TestDelete(t *testing.T) {
	testkit.CheckService()
	defer func() {
		testProd, err := testkit.CreateTestProduct()
		if err != nil {
			t.Log("failed to create the test product:", err)
			t.Fail()
			return
		}

		res, err := testkit.DeleteProduct(testProd.GUID)
		if err != nil {
			t.Logf("delete test encountered an error while deleting the product:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, res.Body, err.Error())
			t.Fail()
			return
		}

		if !testkit.CheckStatusCode(res.StatusCode, expectedDeleteCodes) {
			t.Logf("delete test did not return an expected status code:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, res.Body, err.Error())
			t.Fail()
			return
		}

		// If we can unmarshal the returned JSON then we know it was removed successfully
		_, err = testkit.UnmarshalToProduct(res)
		if err != nil {
			t.Logf("delete test did not return an expected JSON product payload or failed to unmarshal:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, res.Body, err.Error())
			t.Fail()
			return
		}
	}()
	defer testkit.KillService()
}
