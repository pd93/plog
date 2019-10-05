package main

import (
	"errors"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := fileExample(); err != nil {
		log.Fatal(err)
	}
}

func fileExample() (err error) {

	// Open a JSON file
	jsonWriter, err := log.NewFile("log.json")
	if err != nil {
		return err
	}
	defer jsonWriter.Close()

	// Open a CSV file
	csvWriter, err := log.NewFile("log.csv")
	if err != nil {
		return err
	}
	defer csvWriter.Close()

	// Create some loggers
	log.AddLogger("std", log.NewLogger(os.Stdout))
	log.AddLogger("json", log.NewJSONFileLogger(jsonWriter))
	log.AddLogger("csv", log.NewCSVFileLogger(csvWriter))

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
	stdLogger.SetLogLevel(log.TraceLevel)
	stdLogger.SetTimestampFormat("Mon Jan 2 15:04:05 -0700 UTC 2006")

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
