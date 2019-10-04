package plog

import (
	"fmt"
	"strconv"
	"strings"
)

// LogVariables contains a slice of variables that need to be logged
type LogVariables []interface{}

// text will stringify the variables into a readable format and color it if necessary
func (logVariables LogVariables) text() string {

	// Return an empty string if there are no variables
	if len(logVariables) == 0 {
		return ""
	}

	variables := make([]string, len(logVariables))

	// Loop through the variables and format them
	for i, logVariable := range logVariables {
		variables[i] = fmt.Sprintf("%v", logVariable)
	}

	return strings.Join(variables, " ")
}

// json will stringify the variables into a quoted string and color it if necessary
func (logVariables LogVariables) json() string {

	// Return an empty string if there are no variables
	if len(logVariables) == 0 {
		return ""
	}

	variables := make([]string, len(logVariables))

	// Loop through the variables and format them
	for i, logVariable := range logVariables {

		// TODO: We should create a JSON object rather than quoting it as a string
		variables[i] = strconv.Quote(fmt.Sprintf("%v", logVariable))
	}

	return fmt.Sprintf(`[ %s ]`, strings.Join(variables, ", "))
}

// csv will stringify the variables and color it if necessary
func (logVariables LogVariables) csv() string {

	// CSV looks the same as text, so just call the text method
	return logVariables.text()
}
