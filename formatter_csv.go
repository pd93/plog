package plog

import (
	"fmt"
)

// CSVFormatter will create a CSV string using given log and logger configuration
func CSVFormatter(logger *Logger, log *Log) (string, error) {

	// Render each component of the log
	timestamp := log.timestamp.Format(logger.timestampFormat)
	message := log.variables.CSV()
	logLevel := log.logLevel.CSV(logger.colorLogging, logger.logLevelColorMap)
	tags := log.tags.CSV(logger.colorLogging, logger.tagColorMap)

	// Set the output
	return fmt.Sprintf(`%s,%s,%s,%s`, timestamp, logLevel, message, tags), nil
}
