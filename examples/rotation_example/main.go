package main

import (
	"os"

	log "github.com/pd93/plog"
	"github.com/pd93/plog/sequencers"
	"github.com/pd93/plog/writers"
)

func main() {
	if err := RotationExample(); err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}
}

// RotationExample prints logs to a rotating log file.
// This means that each time a log file reaches the rotation criteria,
// the file is closed and a new one is opened in its place.
// You can make a log file rotate by providing a sequencer as a function option.
// This can be set during creation of by using the `file.Options()` method.
// Depending on the sequencer, you may need to provide a formattable string.
// ie. one that can be formatted by `fmt.Sprintf` by including % tags.
// You will need to check the documentation for your specific sequencer.
// If you want to provide your own file rotation sequence,
// you can take a look at the `sequencer_example`.
func RotationExample() (err error) {

	// Open a text file
	rotatingTextFile, err := log.NewFile("./logs/log-%s.txt",
		log.WithSequencer(sequencers.DateTime),
		log.WithMaxFileSize(1024*1024), // 1 MB
	)
	if err != nil {
		return err
	}
	defer rotatingTextFile.Close()

	// Open a JSON file
	rotatingJSONFile, err := log.NewFile("./logs/log-%03d.json",
		log.WithWriter(writers.JSON),
		log.WithSequencer(sequencers.Increment),
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
