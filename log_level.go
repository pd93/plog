package plog

import "strconv"

// LogLevel dictates when a logged message should be displayed or recorded
type LogLevel int

const (
	// None will stop all logs being printed
	None LogLevel = iota - 4 // Start at -4 so that InfoLevel is at 0 (default)
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

func (logLevel LogLevel) String() string {
	switch logLevel {
	case None:
		return "NONE"
	case FatalLevel:
		return "FATAL"
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	case TraceLevel:
		return "TRACE"
	default:
		return ""
	}
}

// text will stringify the log level into a readable format
func (logLevel LogLevel) text(colorLogging bool, logLevelColorMap LogLevelColorMap) string {

	// Check if color logging is enabled and whether there is a color for this log level in the map
	if attributes, ok := logLevelColorMap[logLevel]; colorLogging && ok {
		return color(logLevel.String(), attributes...)
	}

	return logLevel.String()
}

// json will stringify the log level into a quoted string and color it if necessary
func (logLevel LogLevel) json(colorLogging bool, logLevelColorMap LogLevelColorMap) string {

	// Check if color logging is enabled and whether there is a color for this log level in the map
	if attributes, ok := logLevelColorMap[logLevel]; ok && colorLogging {
		return color(strconv.Quote(logLevel.String()), attributes...)
	}

	return strconv.Quote(logLevel.String())
}

// csv will stringify the log level and color it if necessary
func (logLevel LogLevel) csv(colorLogging bool, logLevelColorMap LogLevelColorMap) string {

	// CSV looks the same as text, so just call the text method
	return logLevel.text(colorLogging, logLevelColorMap)
}
