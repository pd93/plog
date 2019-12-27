package main

import (
	"errors"
	"os"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := fileExample(); err != nil {
		log.Fatal(err)
	}
}

func fileExample() (err error) {

	// Open a text file
	textFile, err := log.NewTextFile("./logs/log.txt")
	if err != nil {
		return err
	}
	defer textFile.Close()

	// Open a JSON file
	jsonFile, err := log.NewJSONFile("./logs/log.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Open a CSV file
	csvFile, err := log.NewCSVFile("./logs/log.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// Create some loggers
	log.AddLogger("std", log.NewLogger(os.Stdout))
	log.AddLogger("text", log.NewTextFileLogger(textFile))
	log.AddLogger("json", log.NewJSONFileLogger(jsonFile))
	log.AddLogger("csv", log.NewCSVFileLogger(csvFile))

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
