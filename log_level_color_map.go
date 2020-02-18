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

	// Apply the custom options
	logLevelColorMap.Options(opts...)

	return
}

//
// Functional Options
//

// WithLogLevelColorMapping will return a function that sets a color mapping in a log level color map
func WithLogLevelColorMapping(logLevel LogLevel, attributes ...Attribute) LogLevelColorMapping {
	return func(logLevelColorMap LogLevelColorMap) {
		logLevelColorMap[logLevel] = attributes
	}
}

//
// Options Setter
//

// Options will apply the given options to the log level color map
// Any number of functional options can be passed to this method
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options
func (logLevelColorMap LogLevelColorMap) Options(opts ...LogLevelColorMapping) {
	for _, opt := range opts {
		opt(logLevelColorMap)
	}
}

//
// Getters
//

// Get will return the list of text atrributes for a log level
func (logLevelColorMap LogLevelColorMap) Get(logLevel LogLevel) []Attribute {
	return logLevelColorMap[logLevel]
}
