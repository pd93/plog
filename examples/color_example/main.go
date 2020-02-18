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

	// Create a log level color map
	logLevelColorMap := log.NewLogLevelColorMap(
		log.WithLogLevelColorMapping(log.FatalLevel, log.BgMagenta),
		log.WithLogLevelColorMapping(log.ErrorLevel, log.BgRed),
	)

	// Create a tag color map
	tagColorMap := log.NewTagColorMap(
		log.WithTagColorMapping("tag1", log.BgMagenta),
		log.WithTagColorMapping("tag2", log.BgRed),
	)

	// Create a logger that logs to stdout
	// Log level and tag colors can be set during logger creation using functional options
	log.AddLogger("std", log.NewLogger(
		log.WithLogLevelColorMap(logLevelColorMap),
		log.WithTagColorMap(tagColorMap),
	))

	// Write to all loggers
	log.Fatal(errors.New("Fatal log"))
	log.TError(log.Tags{"tag1"}, errors.New("Error log"))
	log.TInfo(log.Tags{"tag1", "tag2"}, "Info log")
	log.TWarn(log.Tags{"tag2", "tag3"}, "Warn log")

	// Retrieve our standard logger
	stdLogger := log.GetLogger("std")

	// Change the text attributes for each log level / tag individually
	stdLogger.LogLevelColorMap().Options(
		log.WithLogLevelColorMapping(log.FatalLevel, log.BgMagenta),
	)
	stdLogger.TagColorMap().Options(
		log.WithTagColorMapping("tag2", log.BgMagenta),
	)

	// Print a fatal log with our new color settings
	log.TFatal(log.Tags{"tag2"}, errors.New("Fatal log"))

	// OR completely change all the colors!
	// You can change other text attributes too, such as bold, underline and italics

	// Log level colors
	stdLogger.LogLevelColorMap().Options(
		log.WithLogLevelColorMapping(log.FatalLevel, log.FgRed, log.Underline),
		log.WithLogLevelColorMapping(log.ErrorLevel, log.FgRed),
		log.WithLogLevelColorMapping(log.WarnLevel, log.FgYellow, log.Faint),
		log.WithLogLevelColorMapping(log.InfoLevel, log.FgWhite),
	)

	// Tag colors
	stdLogger.TagColorMap().Options(
		// No entry for #tag1 so it prints Faint FgWhite (grey) by default
		log.WithTagColorMapping("tag2", log.FgMagenta),
		log.WithTagColorMapping("tag3", log.FgCyan),
	)

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
