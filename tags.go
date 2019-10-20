package plog

// Tags is a slice of Tag
type Tags []Tag

// String will stringify the tag into a readable format and color it if necessary
func (tags Tags) String(colorLogging bool, tagColorMap TagColorMap) (strs []string) {

	strs = make([]string, len(tags))

	// Loop through the tags and format them
	for i, tag := range tags {
		strs[i] = tag.String(colorLogging, tagColorMap)
	}

	return
}
