package main

import (
	"os"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := sequencerExample(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func sequencerExample() (err error) {

	// Open a text file
	rotatingTextFile, err := log.NewFile("./logs/log.txt",
		log.WithSequencer(func(format, prev string) (next string, err error) {
			return format, nil
		}),
	)
	if err != nil {
		return err
	}
	defer rotatingTextFile.Close()

	// Create some loggers
	log.AddLogger("std", log.NewLogger())
	log.AddLogger("text", log.NewTextFileLogger(rotatingTextFile))

	// Write to all loggers a few times
	for i := 0; i < 10; i++ {
		log.Info("Info log")
	}

	return
}
