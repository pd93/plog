package plog

// ColorMap lists which text attributes should be used for each log level
type ColorMap map[LogLevel][]Attribute

//
// Constructors
//

// NewColorMap creates and returns an instance of ColorMap with the default values
func NewColorMap() (colorMap ColorMap) {
	return ColorMap{
		FatalLevel: []Attribute{FgBlack, BgRed},
		ErrorLevel: []Attribute{FgRed},
		WarnLevel:  []Attribute{FgYellow},
		InfoLevel:  []Attribute{FgGreen},
		DebugLevel: []Attribute{FgCyan},
		TraceLevel: []Attribute{FgBlue},
	}
}

//
// Getters
//

// Get will return the list of text atrributes for a log level
func (colorMap ColorMap) Get(logLevel LogLevel) []Attribute {
	return colorMap[logLevel]
}

//
// Setters
//

// Set will assign a list of text attributes to a log level
func (colorMap ColorMap) Set(logLevel LogLevel, attributes ...Attribute) {
	colorMap[logLevel] = attributes
}
