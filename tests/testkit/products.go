package testkit

import (
	"bytes"
	"encoding/json"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/config"
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	uuid2 "github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

//CreateTestProduct creates a product intended for testkit purposes
func CreateTestProduct() (*models.ProductJSON, error) {
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var testProd models.ProductJSON
	err = json.Unmarshal(body, &testProd)
	if err != nil {
		return nil, err
	}

	return &testProd, nil
}

//CreateTestProductCollection creates a product intended for testkit purposes
func CreateTestProductCollection() (*models.ModelJSONCollection, error) {
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var testProducts models.ModelJSONCollection
	err = json.Unmarshal(body, &testProducts)
	if err != nil {
		return nil, err
	}

	return &testProducts, nil
}

// DeleteProduct will delete the product by its guid
func DeleteProduct(guid string) (*http.Response, error) {
	url := config.Host + "/products/" + guid
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println("[TEST-util] Deleted test product: ", guid)
	return res, err
}

// UnmarshalToProduct will unmarshal a http response that is expected to contain a product JSON payload
func UnmarshalToProduct(res *http.Response) (*models.ProductJSON, error) {
	var product models.ProductJSON
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UnmarshalToCollection will unmarshal a http response that is expected to contain a product collection JSON payload
func UnmarshalToCollection(res *http.Response) (*models.ModelJSONCollection, error) {
	var products models.ModelJSONCollection
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}
	return &products, nil
}
