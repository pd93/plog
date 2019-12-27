package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := writerExample(); err != nil {
		log.Fatal(err)
	}
}

func writerExample() (err error) {

	// Open a text file
	textFile, err := log.NewTextFile("./logs/log.txt")
	if err != nil {
		return err
	}
	defer textFile.Close()

	// Create some loggers
	log.AddLogger("std", log.NewLogger(os.Stdout))
	log.AddLogger("text", log.NewTextFileLogger(textFile))

	// Set the writer
	textFile.SetWriter(func(file *os.File, p []byte) (n int, err error) {

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
	})

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	return
}
