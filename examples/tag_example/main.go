package main

import (
	"errors"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := tagExample(); err != nil {
		log.Fatal(err)
	}
}

func tagExample() (err error) {

	// Create some loggers
	log.AddLogger("std", log.NewLogger())

	// Write to all loggers
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")

	// Write to all loggers again, but with tags this time
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

	// Change the tag colors!
	// You can change other text attributes too, such as bold, underline and italics
	stdLogger.SetTagColorMap(log.TagColorMap{
		// No entry for #tag1 so it prints Faint FgWhite (grey) by default
		"tag2": []log.Attribute{log.FgMagenta},
		"tag3": []log.Attribute{log.FgCyan},
	})

	// Write to all loggers one more time
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
