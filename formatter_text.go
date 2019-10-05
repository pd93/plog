package plog

import (
	"fmt"
)

// TextFormatter will create a Text string using given log and logger configuration
func TextFormatter(logger *Logger, log *Log) (string, error) {

	// Render each component of the log
	timestamp := log.timestamp.Format(logger.timestampFormat)
	message := log.variables.text()
	logLevel := log.logLevel.text(logger.colorLogging, logger.logLevelColorMap)
	tags := log.tags.text(logger.colorLogging, logger.tagColorMap)

	// Set the output
	return fmt.Sprintf("%s [%s] %s %s", timestamp, logLevel, message, tags), nil
}
