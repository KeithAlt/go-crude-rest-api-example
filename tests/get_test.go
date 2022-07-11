package tests

import (
	"bytes"
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/tests/testkit"
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
func TestAllGet(t *testing.T) {
	TestGetAll(t)
	TestGet(t)
}

// TestGetInParallel will run all of our tests in parallel
func TestGetInParallel(t *testing.T) {
	t.Run("TEST GET all Routine", TestGetAll)
	t.Run("TEST GET Routine", TestGet)
}

// TestGetAll tests the get all method of our service
func TestGetAll(t *testing.T) {
	testkit.CheckService()
	defer func() {
		res, err := sendGetRequest("/products")
		if err != nil {
			t.Log("get test HTTP handshake failed: ", err)
			t.Fail()
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Logf("test get all failed to read the response payload body:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		defer res.Body.Close()

		if !testkit.CheckStatusCode(res.StatusCode, expectedGetCodes) {
			t.Logf("test get all request returned an unexpected error code:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

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

	defer testkit.KillService()
}

// TestGet tests the get method of our service
func TestGet(t *testing.T) {
	testkit.CheckService()
	defer func() {
		testProduct, err := testkit.CreateTestProduct()
		if err != nil || testProduct.GUID == "" {
			t.Log("failed to create the test product: ", err, testProduct)
			t.Fail()
			return
		}

		res, err := sendGetRequest("/product/" + testProduct.GUID)
		if err != nil {
			t.Log("get test HTTP handshake failed: ", err)
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

		if !testkit.CheckStatusCode(res.StatusCode, expectedGetCodes) {
			t.Logf("get test request returned an unexpected error code:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		var product models.ProductJSON
		err = json.Unmarshal(body, &product)
		if err != nil {
			t.Logf("get test failed to unmarshal the expected payload:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		_, err = testkit.DeleteProduct(product.GUID)
		if err != nil {
			t.Logf("get test passed but we failed to delete the test product:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}
	}()
	defer testkit.KillService()
}

// sendGetRequest sends a get request with required auth
func sendGetRequest(arg string) (*http.Response, error) {
	var jsonResponse []byte
	req, err := http.NewRequest("GET", config.Host+arg, bytes.NewBuffer(jsonResponse))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;	charset=UTF-8")
	req.Header.Set("Secret", config.Secret)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return res, err
	}
	return res, nil
}
