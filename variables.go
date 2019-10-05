package plog

import (
	"fmt"
	"strconv"
	"strings"
)

// Variables contains a slice of variables that need to be logged
type Variables []interface{}

// Text will stringify the variables into a readable format and color it if necessary
func (variables Variables) Text() string {

	// Return an empty string if there are no variables
	if len(variables) == 0 {
		return ""
	}

	strVariables := make([]string, len(variables))

	// Loop through the variables and format them
	for i, variable := range variables {
		strVariables[i] = fmt.Sprintf("%v", variable)
	}

	return strings.Join(strVariables, " ")
}

// JSON will stringify the variables into a quoted string and color it if necessary
func (variables Variables) JSON() string {

	// Return an empty string if there are no variables
	if len(variables) == 0 {
		return ""
	}

	strVariables := make([]string, len(variables))

	// Loop through the variables and format them
	for i, variable := range variables {

		// TODO: We should create a JSON object rather than quoting it as a string
		strVariables[i] = strconv.Quote(fmt.Sprintf("%v", variable))
	}

	return fmt.Sprintf(`[ %s ]`, strings.Join(strVariables, ", "))
}

// CSV will stringify the variables and color it if necessary
func (variables Variables) CSV() string {

	// CSV looks the same as text, so just call the text method
	return variables.Text()
}
