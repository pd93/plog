package plog

// LogLevelColorMap lists which text attributes should be used for each log level
type LogLevelColorMap map[LogLevel][]Attribute

// A LogLevelColorMapping is a function that sets a color mapping on a given log level color map
type LogLevelColorMapping func(logLevelColorMap LogLevelColorMap)

//
// Constructors
//

// NewLogLevelColorMap creates and returns an instance of LogLevelColorMap with the default values
func NewLogLevelColorMap(opts ...LogLevelColorMapping) (logLevelColorMap LogLevelColorMap) {

	// Create a default log level color map
	logLevelColorMap = LogLevelColorMap{
		FatalLevel: []Attribute{FgBlack, BgRed},
		ErrorLevel: []Attribute{FgRed},
		WarnLevel:  []Attribute{FgYellow},
		InfoLevel:  []Attribute{FgGreen},
		DebugLevel: []Attribute{FgCyan},
		TraceLevel: []Attribute{FgBlue},
	}

	// Loop through each option and call the functional option
	for _, opt := range opts {
		opt(logLevelColorMap)
	}

	return
}

//
// Functional Options
//

// WithLogLevelColorMapping will return a function that sets a color mapping in a log level color map
func WithLogLevelColorMapping(logLevel LogLevel, attributes ...Attribute) LogLevelColorMapping {
	return func(logLevelColorMap LogLevelColorMap) {
		logLevelColorMap.Set(logLevel, attributes...)
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
