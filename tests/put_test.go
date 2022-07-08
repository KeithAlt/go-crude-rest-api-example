package tests

import (
	"net/http"
	"testing"
)

// expected response codes
var expectedPutCodes = []int{
	http.StatusOK,
	http.StatusAccepted,
	http.StatusConflict,
}

// TestPut will run all of our tests for the route
func TestPut(t *testing.T) {
	// TODO test all routes
}

// TestPutStatusCodes tests to ensure the response code is what we expect it to be
func TestPutStatusCodes(t *testing.T) {
	checkService()
	defer func() {
		// TODO implement ...
	}()
	defer killService()
}

// TestPutResponse tests to ensure the response code is what we expect it to be
func TestPutResponse(t *testing.T) {
	checkService()
	defer func() {
		// TODO implement ...
	}()
	defer killService()
}
