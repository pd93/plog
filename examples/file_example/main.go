package main

import (
	"errors"
	"os"

	log "gopkg.in/pd93/plog.v0"
	"gopkg.in/pd93/plog.v0/writers"
)

func main() {
	if err := FileExample(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// FileExample prints log messages to stdout AND two files.
// It does this by setting up multiple loggers, each with their own configuration.
// To create a file logger, it is recommended that you use the `plog.File` wrapper.
// This allows you to take advantage of PLogs file handling features.
// PLog also provides pre-made constructors for creating file loggers in text, JSON and CSV formats.
// These constructors disable colored logging and set up a custom writer for handling those specific file types.
// If you need custom file formats, consider implementing your own formatter and/or writer.
// Take a look at the `formatter_example` and the `writer_example`s respectively to see how this is done.
func FileExample() (err error) {

	// Open a text file
	textFile, err := log.NewFile("./logs/log.txt")
	if err != nil {
		return err
	}
	defer textFile.Close()

	// Open a JSON file
	jsonFile, err := log.NewFile("./logs/log.json",
		log.WithWriter(writers.JSON),
	)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Open a CSV file
	csvFile, err := log.NewFile("./logs/log.csv",
		log.WithWriter(writers.CSV),
	)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// Create some loggers
	log.AddLogger("std", log.NewLogger())
	log.AddLogger("text", log.NewTextFileLogger(textFile))
	log.AddLogger("json", log.NewJSONFileLogger(jsonFile))
	log.AddLogger("csv", log.NewCSVFileLogger(csvFile))

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

	// Write to a specific logger
	stdLogger.Info("Special log that only prints to stdout")

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	return
}
