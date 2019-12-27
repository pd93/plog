package plog

import (
	"testing"
	"time"
)

func TestCSVFormatter(t *testing.T) {

	// Expected output
	const expected = `2006-01-02T15:04:05Z,INFO,Test string 123 4.5 true,tag1:tag2`

	// Test log
	var log = &Log{
		logLevel:  InfoLevel,
		variables: []interface{}{"Test string", 123, 4.5, true},
		timestamp: time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC),
		tags:      Tags{"tag1", "tag2"},
	}

	// Call the function
	output, err := CSVFormatter(
		log.timestamp.Format(time.RFC3339),
		log.logLevel.String(false, NewLogLevelColorMap()),
		log.variables,
		log.tags.String(false, NewTagColorMap()),
	)
	if err != nil {
		t.Error(err)
	}

	// Check if the output is correct
	if output != expected {
		t.Errorf("Incorrect output.\n\tExpected: '%s'\n\tReceived: '%s'", expected, output)
	}
}
