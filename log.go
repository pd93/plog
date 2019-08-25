package plog

import (
	"fmt"
	"time"
)

// A Log holds a single message along with its log level and timestamp
type Log struct {
	LogLevel  LogLevel      `json:"logLevel"`
	Variables []interface{} `json:"message"`
	Timestamp time.Time     `json:"timestamp"`
}

// NewLog creates a new instance of log and populates it with a log level and a message
// A timestamp is also generated and stored
func NewLog(logLevel LogLevel, variables ...interface{}) (log *Log) {
	return &Log{
		LogLevel:  logLevel,
		Variables: variables,
		Timestamp: time.Now(),
	}
}

// NewLogf creates a new instance of log and populates it with a log level and a formatted message
// You can send any number of variables to this function and they will be printed according to the format specified in the message
// A timestamp is also generated and stored
func NewLogf(level LogLevel, message string, variables ...interface{}) (log *Log) {
	return NewLog(level, fmt.Sprintf(message, variables...))
}
