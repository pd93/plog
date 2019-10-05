package plog

import (
	"fmt"
	"strings"
)

// JSONFormatter will create a JSON string using given log and logger configuration
func JSONFormatter(logger *Logger, log *Log) (string, error) {

	var fields []string

	// Timestamp
	if timestamp := log.timestamp.Format(logger.timestampFormat); timestamp != "" {
		fields = append(fields, fmt.Sprintf(`"timestamp": "%s"`, timestamp))
	}

	// Log level
	if logLevel := log.logLevel.json(logger.colorLogging, logger.logLevelColorMap); logLevel != "" {
		fields = append(fields, fmt.Sprintf(`"logLevel": %s`, logLevel))
	}

	// Message
	if message := log.variables.json(); message != "" {
		fields = append(fields, fmt.Sprintf(`"message": %s`, message))
	}

	// Tags
	if tags := log.tags.json(logger.colorLogging, logger.tagColorMap); tags != "" {
		fields = append(fields, fmt.Sprintf(`"tags": %s`, tags))
	}

	// Set the output
	return fmt.Sprintf("{ %s }", strings.Join(fields, ", ")), nil
}
