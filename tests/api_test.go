package tests

import (
	"testing"
)

// TestRunAll tests all of our routes
func TestRunAll(t *testing.T) {
	TestDelete(t)
	TestGet(t)
	TestPost(t)
	TestPut(t)
}

// TestRunAllInParallel will run all of our tests in parallel
func TestRunAllInParallel(t *testing.T) {
	TestDeleteInParallel(t)
	TestGetInParallel(t)
	TestPostAllInParallel(t)
	TestPutInParallel(t)
}
