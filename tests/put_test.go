package tests

import (
	"bytes"
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/tests/testkit"
	uuid2 "github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"testing"
)

// expected response codes
var expectedPutCodes = []int{
	http.StatusOK,
	http.StatusAccepted,
	http.StatusConflict,
}

// TestAllPut will run all of our tests for the route
func TestAllPut(t *testing.T) {
	TestPut(t)
}

// TestPutInParallel will run all of our tests in parallel
func TestPutInParallel(t *testing.T) {
	t.Run("Test Put Code", TestPut)
}

// TestPut tests to ensure the response code is what we expect it to be
func TestPut(t *testing.T) {
	testkit.CheckService()
	defer func() {
		testProd, err := testkit.CreateTestProduct()
		if err != nil {
			t.Log("failed to create the test product:", err)
			t.Fail()
			return
		}

		// Product data we intend to update to the pre-existing product
		newProductName := uuid2.NewString()
		newProductDesc := uuid2.NewString()
		updatedProduct := models.ProductJSON{
			Name:        newProductName,
			Description: newProductDesc,
		}

		res, err := sendPutRequest(testProd.GUID, &updatedProduct)
		if err != nil {
			t.Logf("test put failed to send the request:\n- StatusCode = %v\n- Response = %v\n", res.Status, res.Body)
			t.Fail()
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Logf("test put failed to read the response body:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
		}

		defer res.Body.Close()

		if !testkit.CheckStatusCode(res.StatusCode, expectedPutCodes) {
			t.Logf("test put request returned an unexpected error code:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		var product models.ProductJSON
		err = json.Unmarshal(body, &product)
		if err != nil {
			t.Logf("test put failed to unmarshal the response payload:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, string(body), err.Error())
			t.Fail()
			return
		}

		if product.Name != newProductName || product.Name != newProductDesc {
			t.Logf("test put failed to update data:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, res.Body, err.Error())
			t.Fail()
			return
		}

		res, err = testkit.DeleteProduct(testProd.GUID)
		if err != nil {
			t.Logf("test put encountered an error while deleting the product:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, res.Body, err.Error())
			t.Fail()
			return
		}
	}()
	defer testkit.KillService()
}

// sendPutRequest sends a post request for a single product
func sendPutRequest(id string, newProduct *models.ProductJSON) (*http.Response, error) {
	var jsonData []byte
	jsonData, _ = json.Marshal(newProduct)
	req, err := http.NewRequest("PUT", config.Host+"/products/"+id, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json;	charset=UTF-8")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
