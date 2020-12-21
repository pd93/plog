package main

import (
	"os"

	log "gopkg.in/pd93/plog.v0"
	"gopkg.in/pd93/plog.v0/formatters"
)

func main() {
	if err := GlobalLoggingExample(); err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}
}

// GlobalLoggingExample shows how to enable/disable global logging.
// This flag will enable you to stop a logger from being automatically
// logged to by the global logger map. This can be useful if you want
// a specific logging format for one off logs.
func GlobalLoggingExample() (err error) {

	// Open a text file
	textFile, err := log.NewFile("./logs/log.txt")
	if err != nil {
		return err
	}
	defer textFile.Close()

	// Create some loggers
	log.AddLogger("std", log.NewLogger())
	log.AddLogger("plain", log.NewLogger(
		log.WithGlobalLogging(false),
		log.WithFormatter(formatters.Plain),
	))

	// Write to all loggers, but do not end the line
	// The plain logger will not be written to as global logging is off
	log.Infof("Info log")

	// Append some text using the plain logger
	log.GetLogger("plain").Infof(" | Appended text")

	// Make sure you print a new line before continuing or the next message will stay on the same line
	log.GetLogger("plain").Info()

	// Another regular log message
	log.Info("Info log 2")

	//
	// Do not do this
	//

	// If we enabled global logging on our plain logger, what would happen?
	log.GetLogger("plain").Options(
		log.WithGlobalLogging(true),
	)

	// Now, every message we send to the global logger is printed to stdout twice
	// Once with formatting, and once with no formatting
	log.Info("Info log 3 - This message is duplicated")

	return
}
