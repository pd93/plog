package formatters

import (
	"fmt"
	"strings"
)

// CSV will create a CSV string using given log and logger configuration
func CSV(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {

	strVariables := make([]string, len(variables))

	// Loop through the variables and format them
	for i, variable := range variables {
		strVariables[i] = fmt.Sprintf("%v", variable)
	}

	return fmt.Sprintf(`%s,%s,%s,%s`, timestamp, logLevel, strings.Join(strVariables, " "), strings.Join(tags, ":")), nil
}
