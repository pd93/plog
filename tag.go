package plog

import (
	"fmt"
	"strconv"
	"strings"
)

// A Tag is a metadata string that can be assigned to any log message
type Tag string

// Tags is a slice of Tag
type Tags []Tag

// Text will stringify the tags into a readable format
func (tags Tags) Text() string {

	strSlice := make([]string, len(tags))

	for i, tag := range tags {
		strSlice[i] = fmt.Sprintf("#%s", tag)
	}

	return strings.Join(strSlice, " ")
}

// JSON will stringify the tags into a JSON array
func (tags Tags) JSON() string {

	strSlice := make([]string, len(tags))

	for i, tag := range tags {
		strSlice[i] = strconv.Quote(string(tag))
	}

	return fmt.Sprintf(`[ %s ]`, strings.Join(strSlice, ", "))
}

// CSV will stringify the tags into a CSV compatible string
func (tags Tags) CSV() string {

	strSlice := make([]string, len(tags))

	for i, tag := range tags {
		strSlice[i] = string(tag)
	}

	return strings.Join(strSlice, ":")
}
