package main

import (
	"os"

	log "gopkg.in/pd93/plog.v0"
	"gopkg.in/pd93/plog.v0/sequencers"
	"gopkg.in/pd93/plog.v0/writers"
)

func main() {
	if err := rotationExample(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func rotationExample() (err error) {

	// Open a text file
	rotatingTextFile, err := log.NewFile("./logs/log-%s.txt",
		log.WithSequencer(sequencers.DateTimeSequencer),
		log.WithMaxFileSize(1024*1024), // 1 MB
	)
	if err != nil {
		return err
	}
	defer rotatingTextFile.Close()

	// Open a JSON file
	rotatingJSONFile, err := log.NewFile("./logs/log-%03d.json",
		log.WithWriter(writers.JSONWriter),
		log.WithSequencer(sequencers.IncrementSequencer),
		log.WithMaxFileSize(1024*1024), // 1 MB
	)
	if err != nil {
		return err
	}
	defer rotatingJSONFile.Close()

	// We can change the maximum log file size at any time
	rotatingTextFile.Options(log.WithMaxFileSize(256))
	rotatingJSONFile.Options(log.WithMaxFileSize(256))

	// Create some loggers
	log.AddLogger("std", log.NewLogger())
	log.AddLogger("text", log.NewTextFileLogger(rotatingTextFile))
	log.AddLogger("json", log.NewJSONFileLogger(rotatingJSONFile))

	// Write to all loggers a few times
	for i := 0; i < 10; i++ {
		log.Info("Info log")
	}

	return
}
