package plog

// A Tag is a metadata string that can be assigned to any log message.
type Tag string

// String will stringify the tag into a readable format and color it if necessary.
func (tag Tag) String(colorLogging bool, tagColorMap TagColorMap) string {

	// Check if color logging is enabled and whether there is a color for this tag in the map
	if attributes, ok := tagColorMap[tag]; colorLogging && ok {
		return Color(string(tag), attributes...)
	}

	// If there is no entry in the map, but color logging is still enabled
	if colorLogging {
		return Color(string(tag), FgWhite, Faint)
	}

	return string(tag)
}
