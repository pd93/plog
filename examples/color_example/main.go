package main

import (
	"errors"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := colorExample(); err != nil {
		log.Fatal(err)
	}
}

func colorExample() (err error) {

	// Create some loggers
	log.AddLogger("std", log.NewLogger())

	// Write to all loggers
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")

	// Retrieve our standard logger
	stdLogger := log.GetLogger("std")

	// Change the text attributes for fatal level logs
	stdLogger.LogLevelColorMap().Set(log.FatalLevel, log.BgMagenta)
	log.Fatal(errors.New("Fatal log"))

	// Change all the colors!
	// You can change other text attributes too, such as bold, underline and italics
	stdLogger.SetLogLevelColorMap(log.LogLevelColorMap{
		log.FatalLevel: []log.Attribute{log.FgRed, log.Underline},
		log.ErrorLevel: []log.Attribute{log.FgRed},
		log.WarnLevel:  []log.Attribute{log.FgYellow, log.Faint},
		log.InfoLevel:  []log.Attribute{log.FgWhite},
	})

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")

	return
}
