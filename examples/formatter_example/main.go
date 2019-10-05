package main

import (
	"errors"
	"fmt"
	"os"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := formatterExample(); err != nil {
		log.Fatal(err)
	}
}

func formatterExample() (err error) {

	// Create a logger
	log.AddLogger("std", log.NewLogger(os.Stdout))

	// Write to all loggers
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	// Change the logger's formatter
	log.GetLogger("std").SetFormatter(func(logger *log.Logger, l *log.Log) (str string, err error) {

		// Render each component of the log
		timestamp := l.Timestamp().Format(logger.TimestampFormat())
		message := l.Variables().Text()
		logLevel := l.LogLevel().Text(logger.ColorLogging(), logger.LogLevelColorMap())
		tags := l.Tags().Text(logger.ColorLogging(), logger.TagColorMap())

		// Set the output
		return fmt.Sprintf("%s - %s - {%s} %s", timestamp, logLevel, tags, message), nil
	})

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	return
}
