package plog

import (
	"fmt"
	"time"
)

// A Log holds a single message along with its log level and timestamp
type Log struct {
	LogLevel  LogLevel
	Variables []interface{}
	Timestamp time.Time
}

// NewLog creates a new instance of log and populates it with a log level and a message
// A timestamp is also generated and stored
func NewLog(logLevel LogLevel, variables ...interface{}) *Log {
	return &Log{
		LogLevel:  logLevel,
		Variables: variables,
		Timestamp: time.Now(),
	}
}

// NewLogf creates a new instance of log and populates it with a log level and a formatted message
// You can send any number of variables to this function and they will be printed according to the format specified
// A timestamp is also generated and stored
func NewLogf(level LogLevel, format string, variables ...interface{}) *Log {
	return NewLog(level, fmt.Sprintf(format, variables...))
}
