package plog

import (
	"fmt"
)

// TextFormatter will create a Text string using given log and logger configuration
func TextFormatter(logger *Logger, l *log) (string, error) {

	// Render each component of the log
	timestamp := l.timestamp.Format(logger.timestampFormat)
	message := l.variables.text()
	logLevel := l.logLevel.text(logger.colorLogging, logger.logLevelColorMap)
	tags := l.tags.text(logger.colorLogging, logger.tagColorMap)

	// Set the output
	return fmt.Sprintf("%s [%s] %s %s", timestamp, logLevel, message, tags), nil
}
