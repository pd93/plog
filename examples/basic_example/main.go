package main

import (
	"errors"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := basicExample(); err != nil {
		log.Fatal(err)
	}
}

func basicExample() (err error) {

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
