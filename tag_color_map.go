package plog

// TagColorMap lists which text attributes should be used for each tag
type TagColorMap map[Tag][]Attribute

// A TagColorMapping is a function that sets a color mapping on a given tag color map
type TagColorMapping func(tagColorMap TagColorMap)

//
// Constructors
//

// NewTagColorMap creates and returns an instance of TagColorMap with the default values
func NewTagColorMap(opts ...TagColorMapping) (tagColorMap TagColorMap) {

	// Create a default tag color map
	tagColorMap = TagColorMap{}

	// Loop through each option and call the functional option
	for _, opt := range opts {
		opt(tagColorMap)
	}

	return
}

//
// Functional Options
//

// WithTagColorMapping will return a function that sets a color mapping in a tag color map
func WithTagColorMapping(tag Tag, attributes ...Attribute) TagColorMapping {
	return func(tagColorMap TagColorMap) {
		tagColorMap.Set(tag, attributes...)
	}
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
