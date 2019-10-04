package plog

import (
	"fmt"
	"strconv"
)

// A Tag is a metadata string that can be assigned to any log message
type Tag string

// text will stringify the tag into a readable format and color it if necessary
func (tag Tag) text(colorLogging bool, tagColorMap TagColorMap) string {

	// Check if color logging is enabled and whether there is a color for this tag in the map
	if attributes, ok := tagColorMap[tag]; colorLogging && ok {
		return color(fmt.Sprintf("#%s", tag), attributes...)
	}

	// If there is no entry in the map, but color logging is still enabled
	if colorLogging {
		return color(fmt.Sprintf("#%s", tag), FgWhite, Faint)
	}

	return fmt.Sprintf("#%s", tag)
}

// json will stringify the tag into a quoted string and color it if necessary
func (tag Tag) json(colorLogging bool, tagColorMap TagColorMap) string {

	// Check if color logging is enabled and whether there is a color for this tag in the map
	if attributes, ok := tagColorMap[tag]; colorLogging && ok {
		return color(strconv.Quote(string(tag)), attributes...)
	}

	// If there is no entry in the map, but color logging is still enabled
	if colorLogging {
		return color(strconv.Quote(string(tag)), FgWhite, Faint)
	}

	return strconv.Quote(string(tag))
}

// csv will stringify the tag and color it if necessary
func (tag Tag) csv(colorLogging bool, tagColorMap TagColorMap) string {

	// Check if color logging is enabled and whether there is a color for this tag in the map
	if attributes, ok := tagColorMap[tag]; colorLogging && ok {
		return color(string(tag), attributes...)
	}

	// If there is no entry in the map, but color logging is still enabled
	if colorLogging {
		return color(string(tag), FgWhite, Faint)
	}

	return string(tag)
}