package plog

import (
	"fmt"
)

// A Logger is a channel for writing logs
type Logger struct {
	name   string
	config *Config
}

// Write will add a log message to the logger
func (logger *Logger) Write(log *Log) (err error) {

	var message string

	// Format the message
	message, err = log.String(logger.config)
	if err != nil {
		return err
	}

	// Check that the output is set
	if logger.config.Output == nil {
		return fmt.Errorf("No output is configured for the logger: %s", logger.name)
	}

	// Print the message to the output writer
	fmt.Fprintln(logger.config.Output, message)

	return
}
