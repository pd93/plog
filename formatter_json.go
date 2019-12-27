package plog

import (
	"fmt"
	"strconv"
	"strings"
)

// JSONFormatter will create a JSON string using given log and logger configuration
func JSONFormatter(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {

	var fields []string

	// Timestamp
	if timestamp != "" {
		fields = append(fields, fmt.Sprintf(`"timestamp": "%s"`, timestamp))
	}

	// Log level
	if logLevel != "" {
		fields = append(fields, fmt.Sprintf(`"logLevel": %s`, strconv.Quote(logLevel)))
	}

	// Message
	if len(variables) > 0 {

		strVariables := make([]string, len(variables))

		// Loop through the variables and format them
		for i, variable := range variables {
			// TODO: We should create a JSON object rather than quoting it as a string
			strVariables[i] = strconv.Quote(fmt.Sprintf("%v", variable))
		}

		fields = append(fields, fmt.Sprintf(`"message": [ %s ]`, strings.Join(strVariables, ", ")))
	}

	// Tags
	if len(tags) > 0 {

		// Loop through the tags and quote them
		for _, tag := range tags {
			tag = strconv.Quote(tag)
		}

		fields = append(fields, fmt.Sprintf(`"tags": [ %s ]`, strings.Join(tags, ", ")))
	}

	// Set the output
	return fmt.Sprintf("{ %s }", strings.Join(fields, ", ")), nil
}
