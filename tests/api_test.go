package tests

import (
	"testing"
)

// RunAllRouteTests tests all of our routes
func RunAllRouteTests(t *testing.T) {
	TestDelete(t)
	TestGetAll(t)
	TestPost(t)
	TestPut(t)
}
