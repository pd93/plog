package plog

import (
	"fmt"
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

// WriteToLogger will write the log message to a specific logger
func (log *Log) Write() {

	// Loop through each logger
	for _, logger := range loggers {

		// Test if the logger has a high enough log level
		if logger.LogLevel >= log.LogLevel {

			// Write to the logger
			if err := logger.Write(log); err != nil {
				panic(err)
			}
		}
	}
}
