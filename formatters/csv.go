package formatters

import (
	"fmt"
	"strings"
)

// CSV will format a log into a comma-separated value (CSV) string.
func CSV(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {

	strVariables := make([]string, len(variables))

	// Loop through the variables and format them
	for i, variable := range variables {
		strVariables[i] = fmt.Sprintf("%v", variable)
	}

	return fmt.Sprintf(`%s,%s,%s,%s`, timestamp, logLevel, strings.Join(strVariables, " "), strings.Join(tags, ":")), nil
}
