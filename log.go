package plog

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// A Log holds a single message along with its log level and timestamp
type Log struct {
	LogLevel  LogLevel      `json:"logLevel"`
	Body      []interface{} `json:"message"`
	Timestamp time.Time     `json:"timestamp"`
}

// NewLog creates a new instance of log and populates it with a log level and a message
// A timestamp is also generated and stored
func NewLog(logLevel LogLevel, params ...interface{}) (log *Log) {
	return &Log{
		LogLevel:  logLevel,
		Body:      params,
		Timestamp: time.Now(),
	}
}

// NewLogf creates a new instance of log and populates it with a log level and a formatted message
// You can send any number of parameters to this function and they will be printing according to the format specified in the message
// A timestamp is also generated and stored
func NewLogf(level LogLevel, message string, params ...interface{}) (log *Log) {
	return NewLog(level, fmt.Sprintf(message, params...))
}

// String will convert a log into a string using the provided config
func (log *Log) String(config *Config) (message string, err error) {

	switch config.LogFormat {

	// Text logger
	case TextFormat:

		timestamp := log.Timestamp.Format(config.TimestampFormat)
		level := strings.ToUpper(log.LogLevel.String())

		// Print the message
		if config.ColorLogging {

			switch log.LogLevel {
			case ErrorLevel, FatalLevel:
				level = color(level, FgRed)
			case WarnLevel:
				level = color(level, FgYellow)
			case InfoLevel:
				level = color(level, FgGreen)
			case DebugLevel:
				level = color(level, FgCyan)
			case TraceLevel:
				level = color(level, FgBlue)
			}
		}

		message = fmt.Sprintf("%s [%s]", timestamp, level)

		for _, item := range log.Body {
			message += fmt.Sprintf(" %v", item)
		}

	// JSON logger
	case JSONFormat:
		var b []byte
		b, err = json.Marshal(log)
		message = string(b)
	}

	return
}

// WriteToLogger will write the log message to a specific logger
func (log *Log) WriteToLogger(logger *Logger) {

	var err error

	// Test if this logger has a high enough log level
	if logger.config.LogLevel >= log.LogLevel {
		if err = logger.Write(log); err != nil {
			panic(err)
		}
	}
}

// WriteToAll will write the log message to all loggers
func (log *Log) WriteToAll() {

	// Loop through each logger
	for _, logger := range loggers {
		log.WriteToLogger(logger)
	}
}
