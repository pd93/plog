package plog

// TagColorMap lists which text attributes should be used for each tag
type TagColorMap map[Tag][]Attribute

//
// Constructors
//

// NewTagColorMap creates and returns an instance of TagColorMap with the default values
func NewTagColorMap() (tagColorMap TagColorMap) {
	return TagColorMap{}
}

//
// Getters
//

// Get will return the list of text atrributes for a log level
func (tagColorMap TagColorMap) Get(tag Tag) []Attribute {
	return tagColorMap[tag]
}

//
// Setters
//

// Set will assign a list of text attributes to a log level
func (tagColorMap TagColorMap) Set(tag Tag, attributes ...Attribute) {
	tagColorMap[tag] = attributes
}
