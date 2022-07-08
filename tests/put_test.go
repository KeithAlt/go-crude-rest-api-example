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
	TestPutStatusCodes(t)
	TestPutResponse(t)
}

// TestPutInParallel will run all of our tests in parallel
func TestPutInParallel(t *testing.T) {
	t.Run("Test Put Code (Routine)", TestPutStatusCodes)
	t.Run("Test Put Response (Routine)", TestPutResponse)
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
