package formatters

import (
	"testing"
	"time"
)

func TestPlain(t *testing.T) {

	// Expected output
	const expected = `Test string 123 4.5 true`

	// Test log
	var (
		logLevel  = "INFO"
		variables = []interface{}{"Test string", 123, 4.5, true}
		timestamp = time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC).Format(time.RFC3339)
		tags      = []string{"tag1", "tag2"}
	)

	// Call the function
	output, err := Plain(
		timestamp,
		logLevel,
		variables,
		tags,
	)
	if err != nil {
		t.Error(err)
	}

	// Check if the output is correct
	if output != expected {
		t.Errorf("Incorrect output.\n\tExpected: '%s'\n\tReceived: '%s'", expected, output)
	}
}
