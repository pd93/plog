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

	// If you want to change the color of your tags from the default (grey), see ../color_example/main.go

	return
}
