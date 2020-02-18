package sequencers

import (
	"testing"
)

func TestIncrement(t *testing.T) {

	// Expected output
	const expected = "test-2"

	// Call the function
	next, err := Increment("test-%d", "test-1")
	if err != nil {
		t.Error(err.Error())
	}

	// Check if the output is correct
	if next != expected {
		t.Errorf("Expected: '%s'\nReceived: '%s'", expected, next)
	}
}
