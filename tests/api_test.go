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
