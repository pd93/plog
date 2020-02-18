package sequencers

import (
	"testing"

	"gopkg.in/pd93/plog.v0/mocks"
)

func TestDateTime(t *testing.T) {

	// Swap out Now() with the mock function
	Now = mocks.Now

	// Expected output
	const expected = "test-2006-01-02T15:04:05.000Z"

	// Call the function
	next, err := DateTime("test-%s", "")
	if err != nil {
		t.Error(err.Error())
	}

	// Check if the output is correct
	if next != expected {
		t.Errorf("Expected: '%s'\nReceived: '%s'", expected, next)
	}
}
