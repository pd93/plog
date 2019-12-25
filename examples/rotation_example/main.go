package main

import (
	"os"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := rotationExample(); err != nil {
		log.Fatal(err)
	}
}

func rotationExample() (err error) {

	// Open a text file
	rotatingTextFile, err := log.NewRotatingFile("./logs/log-%s.txt", log.TextWriter, log.DateTimeSequencer)
	if err != nil {
		return err
	}
	defer rotatingTextFile.Close()

	// Open a JSON file
	rotatingJSONFile, err := log.NewRotatingFile("./logs/log-%03d.json", log.JSONWriter, log.IncrementSequencer)
	if err != nil {
		return err
	}
	defer rotatingJSONFile.Close()

	// Set the maximum log file size
	rotatingTextFile.SetMaxFileSize(256)
	rotatingJSONFile.SetMaxFileSize(256)

	// Create some loggers
	log.AddLogger("std", log.NewLogger(os.Stdout))
	log.AddLogger("text", log.NewTextFileLogger(rotatingTextFile))
	log.AddLogger("json", log.NewJSONFileLogger(rotatingJSONFile))

	// Write to all loggers a few times
	for i := 0; i < 10; i++ {
		log.Info("Info log")
	}

	return
}
