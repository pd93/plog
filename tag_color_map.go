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

	// Apply the custom options
	tagColorMap.Options(opts...)

	return
}

//
// Functional Options
//

// WithTagColorMapping will return a function that sets a color mapping in a tag color map
func WithTagColorMapping(tag Tag, attributes ...Attribute) TagColorMapping {
	return func(tagColorMap TagColorMap) {
		tagColorMap[tag] = attributes
	}
}

//
// Options Setter
//

// Options will apply the given options to the tag color map
// Any number of functional options can be passed to this method
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options
func (tagColorMap TagColorMap) Options(opts ...TagColorMapping) {
	for _, opt := range opts {
		opt(tagColorMap)
	}
}

//
// Getters
//

// Get will return the list of text atrributes for a log level
func (tagColorMap TagColorMap) Get(tag Tag) []Attribute {
	return tagColorMap[tag]
}
