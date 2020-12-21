package main

import (
	"errors"
	"os"

	log "github.com/pd93/plog"
)

func main() {
	if err := BasicExample(); err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}
}

// BasicExample creates a stdout logger and prints some random logs.
// It then changes a few basic settings and prints the same logs again.
func BasicExample() (err error) {

	// Create a logger that logs to stdout
	log.AddLogger("std", log.NewLogger())

	// Write to all loggers
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	// Retrieve our standard logger
	stdLogger := log.GetLogger("std")

	// Change some settings
	stdLogger.Options(
		log.WithLogLevel(log.TraceLevel),
		log.WithTimestampFormat("Mon Jan 2 15:04:05 -0700 UTC 2006"),
	)

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	return
}
