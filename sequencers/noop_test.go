package sequencers

import (
	"testing"
)

func TestNoop(t *testing.T) {

	// Expected output
	const expected = "test"

	// Call the function
	next, err := Noop("test", "test")
	if err != nil {
		t.Error(err.Error())
	}

	// Check if the output is correct
	if next != expected {
		t.Errorf("Expected: '%s'\nReceived: '%s'", expected, next)
	}
}
