package plog

import (
	"fmt"
	"time"
)

// A log holds a single message along with its log level and timestamp
type log struct {
	logLevel  LogLevel
	variables []interface{}
	timestamp time.Time
}

// newLog creates a new instance of log and populates it with a log level and a message
// A timestamp is also generated and stored
func newLog(logLevel LogLevel, variables ...interface{}) *log {
	return &log{
		logLevel:  logLevel,
		variables: variables,
		timestamp: time.Now(),
	}
}

// newLogf creates a new instance of log and populates it with a log level and a formatted message
// You can send any number of variables to this function and they will be printed according to the format specified
// A timestamp is also generated and stored
func newLogf(level LogLevel, format string, variables ...interface{}) *log {
	return newLog(level, fmt.Sprintf(format, variables...))
}
