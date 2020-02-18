package main

import (
	"os"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := SequencerExample(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// SequencerExample prints log messages to a file using a custom file rotation sequencer.
// All sequencers must take a file name format and the name of the previous file.
// It is then up to the implementer to decide how this data should be used to contruct the new file name.
// In the example below, we simply return the given file format.
// This means that the file name will be the same on every rotation and will therefore overwrite the last log file.
// NOTE: This is essentially the same as `sequencers.Noop`
func SequencerExample() (err error) {

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
