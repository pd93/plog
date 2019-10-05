package plog

import (
	"fmt"
	"strings"
)

// Tags is a slice of Tag
type Tags []Tag

// Text will stringify the tag into a readable format
func (tags Tags) Text(colorLogging bool, tagColorMap TagColorMap) string {

	// Return an empty string if there are no tags
	if len(tags) == 0 {
		return ""
	}

	strSlice := make([]string, len(tags))

	// Loop through the tags and format them
	for i, tag := range tags {
		strSlice[i] = tag.Text(colorLogging, tagColorMap)
	}

	return strings.Join(strSlice, " ")
}

// JSON will stringify the tags into a JSON array
func (tags Tags) JSON(colorLogging bool, tagColorMap TagColorMap) string {

	// Return an empty string if there are no tags
	if len(tags) == 0 {
		return ""
	}

	strSlice := make([]string, len(tags))

	// Loop through the tags and format them
	for i, tag := range tags {
		strSlice[i] = tag.JSON(colorLogging, tagColorMap)
	}

	return fmt.Sprintf(`[ %s ]`, strings.Join(strSlice, ", "))
}

// CSV will stringify the tags into a CSV compatible string
func (tags Tags) CSV(colorLogging bool, tagColorMap TagColorMap) string {

	// Return an empty string if there are no tags
	if len(tags) == 0 {
		return ""
	}

	strSlice := make([]string, len(tags))

	// Loop through the tags and format them
	for i, tag := range tags {
		strSlice[i] = tag.CSV(colorLogging, tagColorMap)
	}

	return strings.Join(strSlice, ":")
}
