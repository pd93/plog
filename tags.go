package plog

import (
	"fmt"
	"strings"
)

// Tags is a slice of Tag
type Tags []Tag

// text will stringify the tag into a readable format
func (tags Tags) text(colorLogging bool, tagColorMap TagColorMap) string {

	// Return an empty string if there are no tags
	if len(tags) == 0 {
		return ""
	}

	strSlice := make([]string, len(tags))

	// Loop through the tags and format them
	for i, tag := range tags {
		strSlice[i] = tag.text(colorLogging, tagColorMap)
	}

	return strings.Join(strSlice, " ")
}

// json will stringify the tags into a JSON array
func (tags Tags) json(colorLogging bool, tagColorMap TagColorMap) string {

	// Return an empty string if there are no tags
	if len(tags) == 0 {
		return ""
	}

	strSlice := make([]string, len(tags))

	// Loop through the tags and format them
	for i, tag := range tags {
		strSlice[i] = tag.json(colorLogging, tagColorMap)
	}

	return fmt.Sprintf(`[ %s ]`, strings.Join(strSlice, ", "))
}

// csv will stringify the tags into a CSV compatible string
func (tags Tags) csv(colorLogging bool, tagColorMap TagColorMap) string {

	// Return an empty string if there are no tags
	if len(tags) == 0 {
		return ""
	}

	strSlice := make([]string, len(tags))

	// Loop through the tags and format them
	for i, tag := range tags {
		strSlice[i] = tag.csv(colorLogging, tagColorMap)
	}

	return strings.Join(strSlice, ":")
}
