package plog

import (
	"fmt"
	"regexp"
	"strings"
)

// TextFormatter will create a readable string using given log and logger configuration
// TODO: There is probably a better way to put the tag inside the color formatting
// TODO: Should the color formatting happen here?
func TextFormatter(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {

	strVariables := make([]string, len(variables))

	// Loop through the variables and format them
	for i, variable := range variables {
		strVariables[i] = fmt.Sprintf("%v", variable)
	}

	// Create a regular expression to check for color in the tags
	re := regexp.MustCompile("^(\\x1b\\[\\d{1,2}(?:;\\d{1,2})*m)(.*)(\\x1b\\[0m)$")

	// Loop through the tags and format them
	for i := range tags {

		// Find out if the tag has color formatting
		matches := re.FindAllStringSubmatch(tags[i], -1)

		if len(matches) > 0 {
			// Add a '#' INSIDE the color formatting string
			tags[i] = fmt.Sprintf("%s#%s%s", matches[0][1], matches[0][2], matches[0][3])
		} else {
			// Add a '#'
			tags[i] = fmt.Sprintf("#%s", tags[i])
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s [%s] %s %s", timestamp, logLevel, strings.Join(strVariables, " "), strings.Join(tags, " "))), nil
}
