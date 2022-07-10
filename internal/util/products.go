package util

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"reflect"
	"strings"
)

// MergeModelsIntoInterface takes a target model & merges it with a JSON model values
// XXX This works but the reflection pkg is notoriously slow. We can do better!
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
