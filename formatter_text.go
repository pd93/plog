package plog

import (
	"fmt"
	"strings"
)

// TextFormatter will create a readable string using given log and logger configuration
func TextFormatter(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {

	strVariables := make([]string, len(variables))

	// Loop through the variables and format them
	for i, variable := range variables {
		strVariables[i] = fmt.Sprintf("%v", variable)
	}

	// Loop through the tags and format them
	// TODO: Check for color and insert '#' inside color boundary
	for _, tag := range tags {
		tag = fmt.Sprintf("#%s", tag)
	}

	return strings.TrimSpace(fmt.Sprintf("%s [%s] %s %s", timestamp, logLevel, strings.Join(strVariables, " "), tags)), nil
}
