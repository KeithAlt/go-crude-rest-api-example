package util

import (
	"github.com/KeithAlt/go-crude-rest-api-boilerplate/internal/service/models"
	"strings"
)

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
