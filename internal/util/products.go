package util

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"reflect"
	"strings"
)

// MergeModelsIntoInterface takes a target model & merges it with a JSON model interface vars (slowly...)
// TODO remove by 7/22/2022 if marked as unused by IDE. If you're reading this & it's past that date, do it!
func MergeModelsIntoInterface(curModel *models.Product, newModel *models.Product) []interface{} {
	curValue := reflect.ValueOf(curModel)
	newValue := reflect.ValueOf(newModel)
	newModelValues := make([]interface{}, newValue.NumField())
	for i := 0; i < curValue.NumField(); i++ {
		if newValue.Field(i).Interface() == "" {
			newModelValues[i] = curValue.Field(i).Interface()
		} else {
			newModelValues[i] = newValue.Field(i).Interface()
		}
	}
	return newModelValues
}

// IsDuplicateKeyError returns if a returned error is due to a duplicate key constraint conflict
func IsDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

// MergeProductModels merges product models overriding new values of null with current product data
func MergeProductModels(newModel *models.Product, curModel *models.Product) *models.Product {
	mergedProduct := models.Product{
		Name:        getValidFieldByString(newModel.Name, curModel.Name),
		Price:       getValidFieldByFloat32(newModel.Price, curModel.Price),
		Description: getValidFieldByString(newModel.Description, curModel.Description),
		CreatedAt:   curModel.CreatedAt,
		UpdatedAt:   GetTime(),
		GUID:        curModel.GUID,
	}
	return &mergedProduct
}

// getValidFieldByString returns the valid value between two string args
func getValidFieldByString(v1, v2 string) string {
	if v1 == "0" || v2 == "" {
		return v2
	}
	return v1
}

// getValidFieldByFloat32 returns the valid value between two float32 args
func getValidFieldByFloat32(v1, v2 float32) float32 {
	if v1 == 0 {
		return v2
	}
	return v1
}
