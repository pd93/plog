package plog

import (
	"fmt"
	"strings"
)

// JSONFormatter will create a JSON string using given log and logger configuration
func JSONFormatter(logger *Logger, l *log) (string, error) {

	var fields []string

	// Timestamp
	if timestamp := l.timestamp.Format(logger.timestampFormat); timestamp != "" {
		fields = append(fields, fmt.Sprintf(`"timestamp": "%s"`, timestamp))
	}

	// Log level
	if logLevel := l.logLevel.json(logger.colorLogging, logger.logLevelColorMap); logLevel != "" {
		fields = append(fields, fmt.Sprintf(`"logLevel": %s`, logLevel))
	}

	// Message
	if message := l.variables.json(); message != "" {
		fields = append(fields, fmt.Sprintf(`"message": %s`, message))
	}

	// Tags
	if tags := l.tags.json(logger.colorLogging, logger.tagColorMap); tags != "" {
		fields = append(fields, fmt.Sprintf(`"tags": %s`, tags))
	}

	// Set the output
	return fmt.Sprintf("{ %s }", strings.Join(fields, ", ")), nil
}
