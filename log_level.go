package plog

// LogLevel dictates when a logged message should be displayed or recorded.
type LogLevel int

// Available log levels:
const (
	// None will stop all logs being printed
	None LogLevel = iota
	// FatalLevel should only be used to log errors that stop the program from continuing execution
	FatalLevel
	// ErrorLevel should be used to display non-fatal errors
	ErrorLevel
	// WarnLevel should be used when an error isn't appropriate, but the message needs be highlighted
	WarnLevel
	// InfoLevel should be used for logging events and basic information
	InfoLevel
	// DebugLevel should be used for logging variables during debugging
	DebugLevel
	// TraceLevel should be used for extremely detailed logs
	TraceLevel
)

// String will stringify the log level into a readable format and color it if necessary.
func (logLevel LogLevel) String(colorLogging bool, logLevelColorMap LogLevelColorMap) (str string) {

	// Get the log level text
	switch logLevel {
	case None:
		str = "NONE"
	case FatalLevel:
		str = "FATAL"
	case ErrorLevel:
		str = "ERROR"
	case WarnLevel:
		str = "WARN"
	case InfoLevel:
		str = "INFO"
	case DebugLevel:
		str = "DEBUG"
	case TraceLevel:
		str = "TRACE"
	default:
		str = ""
	}

	// Check if color logging is enabled and whether there is a color for this log level in the map
	if attributes, ok := logLevelColorMap[logLevel]; colorLogging && ok {
		return color(str, attributes...)
	}

	return
}
