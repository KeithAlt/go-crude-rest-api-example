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
var expectedPostCodes = []int{
	http.StatusCreated,
	http.StatusOK,
	http.StatusAccepted,
	http.StatusConflict,
}

// TestGetAll will run all of our tests for the route
func TestPostAll(t *testing.T) {
	TestPost(t)
	TestPostCollection(t)
}

// TestPostAllInParallel will run all of our tests in parallel
func TestPostAllInParallel(t *testing.T) {
	t.Run("Test Post (Routine)", TestPost)
	t.Run("Test Post Collection (Routine)", TestPostCollection)
}

// TestPost tests the post method to ensure it returns what we expect it to
func TestPost(t *testing.T) {
	testkit.CheckService()
	defer func() {
		res, err := sendPostRequest()
		if err != nil {
			t.Log("failed to create the test product: ", err)
			t.Fail()
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Logf("test post failed to read the response body:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
		}

		if !testkit.CheckStatusCode(res.StatusCode, expectedPostCodes) {
			t.Logf("test post request returned an unexpected error code:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		defer res.Body.Close()

		var product models.ProductJSON
		err = json.Unmarshal(body, &product)
		if err != nil {
			t.Logf("test post failed to unmarshal the response payload:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, string(body), err.Error())
			t.Fail()
			return
		}

		_, err = testkit.DeleteProduct(product.GUID)
		if err != nil {
			t.Logf("test post passed prior tests but failed to delete the test product:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
		}
	}()
	defer testkit.KillService()
}

// TestPostCollection tests a "bulk post" request response
func TestPostCollection(t *testing.T) {
	testkit.CheckService()
	defer func() {
		res, err := sendPostCollectionRequest()
		if err != nil {
			t.Log()
			t.Fail()
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Log()
			t.Fail()
			return
		}

		if !testkit.CheckStatusCode(res.StatusCode, expectedPostCodes) {
			t.Logf("test post request returned an unexpected error code:\n- StatusCode = %v\n- Response = %v\n", res.Status, string(body))
			t.Fail()
			return
		}

		defer res.Body.Close()

		var testProducts models.ModelJSONCollection
		err = json.Unmarshal(body, &testProducts.Repo)
		if err != nil {
			t.Logf("test post failed to unmarshal the response payload:\n- StatusCode = %v\n- Response = %v\n- Error = %s", res.Status, string(body), err.Error())
			t.Fail()
			return
		}

		for _, p := range testProducts.Repo {
			_, err = testkit.DeleteProduct(p.GUID)
			if err != nil {
				t.Log("test passed but failed to delete products: ", err)
				t.Fail()
			}
		}
	}()
	defer testkit.KillService()
}

// sendPostCollectionRequest sends a request to create a collection of products
func sendPostCollectionRequest() (*http.Response, error) {
	jsonData := []byte(`
		[{
			"name":        "` + uuid2.NewString() + `",
			"price":       500.00,
			"description":  "` + uuid2.NewString() + `"
		},
		{
			"name":        "` + uuid2.NewString() + `",
			"price":       500.00,
			"description":  "` + uuid2.NewString() + `"
		}]`)

	req, err := http.NewRequest("POST", config.Host+"/products", bytes.NewBuffer(jsonData))
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

// sendPostRequest sends a post request for a single product
func sendPostRequest() (*http.Response, error) {
	jsonData := []byte(`{
			"name":        "` + uuid2.NewString() + `",
			"price":       500.00,
			"description":  "` + uuid2.NewString() + `"
	}`)

	req, err := http.NewRequest("POST", config.Host+"/products", bytes.NewBuffer(jsonData))
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
