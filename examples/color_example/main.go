package main

import (
	"errors"
	"os"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := colorExample(); err != nil {
		log.Fatal(err)
	}
}

func colorExample() (err error) {

	// Create some loggers
	log.AddLogger("std", log.NewLogger(os.Stdout))

	// Write to all loggers
	log.Fatal(errors.New("Fatal log"))
	log.TError(
		log.Tags{"tag1"},
		errors.New("Error log"),
	)
	log.TInfo(
		log.Tags{"tag1", "tag2"},
		"Info log",
	)
	log.TWarn(
		log.Tags{"tag2", "tag3"},
		"Warn log",
	)

	// Retrieve our standard logger
	stdLogger := log.GetLogger("std")

	// Change the text attributes for each log level / tag individually
	stdLogger.LogLevelColorMap().Set(log.FatalLevel, log.BgMagenta)
	stdLogger.TagColorMap().Set("tag2", log.BgMagenta)

	// Print a fatal log with our new color settings
	log.Fatal(errors.New("Fatal log"))

	// OR completely change all the colors!
	// You can change other text attributes too, such as bold, underline and italics

	// Log level colors
	stdLogger.SetLogLevelColorMap(log.LogLevelColorMap{
		log.FatalLevel: []log.Attribute{log.FgRed, log.Underline},
		log.ErrorLevel: []log.Attribute{log.FgRed},
		log.WarnLevel:  []log.Attribute{log.FgYellow, log.Faint},
		log.InfoLevel:  []log.Attribute{log.FgWhite},
	})

	// Tag colors
	stdLogger.SetTagColorMap(log.TagColorMap{
		// No entry for #tag1 so it prints Faint FgWhite (grey) by default
		"tag2": []log.Attribute{log.FgMagenta},
		"tag3": []log.Attribute{log.FgCyan},
	})

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.TError(
		log.Tags{"tag1"},
		errors.New("Error log"),
	)
	log.TInfo(
		log.Tags{"tag1", "tag2"},
		"Info log",
	)
	log.TWarn(
		log.Tags{"tag2", "tag3"},
		"Warn log",
	)

	return
}
