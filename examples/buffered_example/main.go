package main

import (
	"errors"
	"fmt"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := plog(); err != nil {
		log.Fatal(err)
	}
}

func plog() (err error) {

	// Open a JSON file with a buffered writer of size 5
	jsonWriter, err := log.NewBufWriter("log.json", 5)
	if err != nil {
		return err
	}
	defer jsonWriter.Close()

	// Create some loggers
	log.AddLogger("json", log.NewJSONFileLogger(jsonWriter))

	// Write to all loggers
	log.Fatal(errors.New("Fatal log"))
	fmt.Printf("Lines in buffer: %d\n", jsonWriter.Size())

	log.Error(errors.New("Error log"))
	fmt.Printf("Lines in buffer: %d\n", jsonWriter.Size())

	log.Warn("Warn log")
	fmt.Printf("Lines in buffer: %d\n", jsonWriter.Size())

	log.Info("Info log")
	fmt.Printf("Lines in buffer: %d\n", jsonWriter.Size())

	log.Debug("Debug log", []string{"Random", "string", "array"})
	fmt.Printf("Lines in buffer: %d\n", jsonWriter.Size())

	// This doesn't get logged yet, because its still in the buffer
	log.Trace("Trace log", true, 42, 3.14159)
	fmt.Printf("Lines in buffer: %d\n", jsonWriter.Size())

	// If necessary, we can manually flush the buffer to the file
	jsonWriter.Flush()
	fmt.Printf("Lines in buffer: %d\n", jsonWriter.Size())

	//
	// As soon as we return, jsonWriter.Close() is called and the file is flushed anyway.
	//

	return
}
