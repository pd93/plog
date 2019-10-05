package plog

import (
	"fmt"
)

// CSVFormatter will create a CSV string using given log and logger configuration
func CSVFormatter(logger *Logger, l *log) (string, error) {

	// Render each component of the log
	timestamp := l.timestamp.Format(logger.timestampFormat)
	message := l.variables.csv()
	logLevel := l.logLevel.csv(logger.colorLogging, logger.logLevelColorMap)
	tags := l.tags.csv(logger.colorLogging, logger.tagColorMap)

	// Set the output
	return fmt.Sprintf(`%s,%s,%s,%s`, timestamp, logLevel, message, tags), nil
}
