package formatters_test

import (
	"testing"
	"time"

	log "gopkg.in/pd93/plog.v0"
	"gopkg.in/pd93/plog.v0/formatters"
)

func TestJSONFormatter(t *testing.T) {

	// Expected output
	const expected = `{"timestamp":"2006-01-02T15:04:05Z","logLevel":"INFO","variables":["Test string",123,4.5,true],"tags":["tag1","tag2"]}`

	// Test log
	var (
		logLevel  = log.InfoLevel.String(false, log.NewLogLevelColorMap())
		variables = []interface{}{"Test string", 123, 4.5, true}
		timestamp = time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC).Format(time.RFC3339)
		tags      = log.Tags{"tag1", "tag2"}.String(false, log.NewTagColorMap())
	)

	// Call the function
	output, err := formatters.JSONFormatter(
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
