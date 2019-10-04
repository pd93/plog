package plog

// LogLevelColorMap lists which text attributes should be used for each log level
type LogLevelColorMap map[LogLevel][]Attribute

//
// Constructors
//

// NewLogLevelColorMap creates and returns an instance of LogLevelColorMap with the default values
func NewLogLevelColorMap() LogLevelColorMap {
	return LogLevelColorMap{
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
func (LogLevelColorMap LogLevelColorMap) Get(logLevel LogLevel) []Attribute {
	return LogLevelColorMap[logLevel]
}

//
// Setters
//

// Set will assign a list of text attributes to a log level
func (LogLevelColorMap LogLevelColorMap) Set(logLevel LogLevel, attributes ...Attribute) {
	LogLevelColorMap[logLevel] = attributes
}
