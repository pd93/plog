package main

import (
	"errors"
	"os"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := TagExample(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// TagExample prints tagged logs to stdout.
// A tag is a way of adding context to a log message.
// For example, you might want to add a tag for the component that logged the message.
// Tags make it easier to filter log messages when you're looking for that pesky error.
// You can apply as many tags to a log message as you like.
// Every logging function has an equivellent "tag logging function".
// e.g. To use a `log.Info` with tags, you should call `log.TInfo`.
func TagExample() (err error) {

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
