package util

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"reflect"
)

// MergeModelsIntoInterface takes a target model & merges it with a JSON model values
// XXX This works but the reflection pkg is notoriously slow. We can do better!
func MergeModelsIntoInterface(curModel *models.Product, newModel *models.ProductJSON) []interface{} {
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
