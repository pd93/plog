package plog

import (
	"fmt"
	"time"
)

// A Log holds a single message along with its log level and timestamp
type Log struct {
	logLevel  LogLevel
	variables Variables
	timestamp time.Time
	tags      Tags
}

// newLog creates a new instance of log and populates it with a log level and a message
// A timestamp is also generated and stored
func newLog(logLevel LogLevel, variables ...interface{}) *Log {
	return &Log{
		logLevel:  logLevel,
		variables: variables,
		timestamp: time.Now(),
	}
}

// newLogf creates a new instance of log and populates it with a log level and a formatted message
// You can send any number of variables to this function and they will be printed according to the format specified
// A timestamp is also generated and stored
func newLogf(level LogLevel, format string, variables ...interface{}) *Log {
	return newLog(level, fmt.Sprintf(format, variables...))
}

// newTLog creates a new instance of log and populates it with a log level, a message and a series of meta-tags
// A timestamp is also generated and stored
func newTLog(logLevel LogLevel, tags Tags, variables ...interface{}) *Log {
	return &Log{
		logLevel:  logLevel,
		variables: variables,
		timestamp: time.Now(),
		tags:      tags,
	}
}

// newTLogf creates a new instance of log and populates it with a log level, a formatted message and a series of meta-tags
// You can send any number of variables to this function and they will be printed according to the format specified
// A timestamp is also generated and stored
func newTLogf(level LogLevel, tags Tags, format string, variables ...interface{}) *Log {
	return newTLog(level, tags, fmt.Sprintf(format, variables...))
}

//
// Getters
//

// LogLevel will return the level at which this log is set to write
func (log *Log) LogLevel() LogLevel {
	return log.logLevel
}

// Variables will return the list of log output variables
func (log *Log) Variables() Variables {
	return log.variables
}

// Timestamp will return the log creation timestamp
func (log *Log) Timestamp() time.Time {
	return log.timestamp
}

// Tags will return the list of tags associated with the log
func (log *Log) Tags() Tags {
	return log.tags
}
