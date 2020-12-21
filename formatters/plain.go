package formatters

import (
	"fmt"
	"strings"
)

// Plain will print a plain text string.
func Plain(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {

	strVariables := make([]string, len(variables))

	// Loop through the variables and format them
	for i, variable := range variables {
		strVariables[i] = fmt.Sprintf("%v", variable)
	}

	return fmt.Sprintf("%s", strings.Join(strVariables, " ")), nil
}
