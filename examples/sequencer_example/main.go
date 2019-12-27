package main

import (
	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := sequencerExample(); err != nil {
		log.Fatal(err)
	}
}

func sequencerExample() (err error) {

	// Open a text file
	rotatingTextFile, err := log.NewRotatingFile("./logs/log-%s.txt", log.TextWriter, log.DateTimeSequencer)
	if err != nil {
		return err
	}
	defer rotatingTextFile.Close()

	// Set provide a custom sequencer function
	rotatingTextFile.SetSequencer(func(format, prev string) (next string, err error) {
		return prev, nil
	})

	// Create some loggers
	log.AddLogger("std", log.NewLogger())
	log.AddLogger("text", log.NewTextFileLogger(rotatingTextFile))

	// Write to all loggers a few times
	for i := 0; i < 10; i++ {
		log.Info("Info log")
	}

	return
}
