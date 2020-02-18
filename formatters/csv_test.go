package formatters_test

import (
	"testing"
	"time"

	log "gopkg.in/pd93/plog.v0"
	"gopkg.in/pd93/plog.v0/formatters"
)

func TestCSV(t *testing.T) {

	// Expected output
	const expected = `2006-01-02T15:04:05Z,INFO,Test string 123 4.5 true,tag1:tag2`

	// Test log
	var (
		logLevel  = log.InfoLevel.String(false, log.NewLogLevelColorMap())
		variables = []interface{}{"Test string", 123, 4.5, true}
		timestamp = time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC).Format(time.RFC3339)
		tags      = log.Tags{"tag1", "tag2"}.String(false, log.NewTagColorMap())
	)

	// Call the function
	output, err := formatters.CSV(
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
