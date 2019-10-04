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
	tags      Tags
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

// newTLog creates a new instance of log and populates it with a log level, a message and a series of meta-tags
// A timestamp is also generated and stored
func newTLog(logLevel LogLevel, tags Tags, variables ...interface{}) *log {
	return &log{
		logLevel:  logLevel,
		variables: variables,
		timestamp: time.Now(),
		tags:      tags,
	}
}

// newLogf creates a new instance of log and populates it with a log level and a formatted message
// You can send any number of variables to this function and they will be printed according to the format specified
// A timestamp is also generated and stored
func newLogf(level LogLevel, format string, variables ...interface{}) *log {
	return newLog(level, fmt.Sprintf(format, variables...))
}

// newTLogf creates a new instance of log and populates it with a log level, a formatted message and a series of meta-tags
// You can send any number of variables to this function and they will be printed according to the format specified
// A timestamp is also generated and stored
func newTLogf(level LogLevel, tags Tags, format string, variables ...interface{}) *log {
	return newLog(level, tags, fmt.Sprintf(format, variables...))
}
