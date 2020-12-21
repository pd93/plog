package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := WriterExample(); err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}
}

// WriterExample prints log messages to stdout and a text file using a custom writer.
// A writer allows the user to write log messages to a file in a custom way.
// This could be to add a header row (e.g. for CSV files) or to surround the log
// messages with curly braces and end logs with commas (e.g. for JSON files).
// JSON and CSV writers are available out of the box, but if you want custom behaviour,
// you can simply implement your own as shown below.
// This custom writer will initialise a file with a 'created' timestamp if the file is empty.
// Otherwise it will write the log to the end of the file as normal.
func WriterExample() (err error) {

	// Open a text file
	textFile, err := log.NewTextFile("./logs/log.txt")
	if err != nil {
		return err
	}
	defer textFile.Close()

	// Create some loggers
	log.AddLogger("std", log.NewLogger())
	log.AddLogger("text", log.NewTextFileLogger(textFile))

	// Set the writer
	textFile.Options(log.WithWriter(func(file *os.File, p []byte) (n int, err error) {

		// Get the file size
		fileInfo, err := file.Stat()
		if err != nil {
			return
		}
		fileSize := fileInfo.Size()

		// If the file is empty, initialise it
		if fileSize == 0 {
			_, err = file.Write([]byte(fmt.Sprintf("--- Created: %s ---\n\n", time.Now().UTC().Format(time.RFC3339))))
			if err != nil {
				return 0, err
			}
		}

		// Write the log message
		if n, err = file.Write(p); err != nil {
			return
		}

		return
	}))

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	return
}
