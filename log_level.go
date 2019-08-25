package plog

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
		return "None"
	case FatalLevel:
		return "Fatal"
	case ErrorLevel:
		return "Error"
	case WarnLevel:
		return "Warn"
	case InfoLevel:
		return "Info"
	case DebugLevel:
		return "Debug"
	case TraceLevel:
		return "Trace"
	default:
		return ""
	}
}
