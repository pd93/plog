package plog

import (
	"testing"
)

type colorTest struct {
	message    string
	attributes []Attribute
	expected   string
}

func TestColor(t *testing.T) {

	tests := []colorTest{
		colorTest{"test", []Attribute{FgRed}, "\x1b[31mtest\x1b[0m"},
		colorTest{"test", []Attribute{FgRed, Underline}, "\x1b[31;4mtest\x1b[0m"},
		colorTest{"test", []Attribute{FgRed, FgBlue}, "\x1b[31;34mtest\x1b[0m"},
	}

	// Loop through the tests
	for i, test := range tests {

		// Call the function
		output := color(test.message, test.attributes...)

		// Check if the output is correct
		if output != test.expected {
			t.Errorf("[%d] Incorrect output. Expected '%s', received '%s'", i, test.expected, output)
		}
	}
}
