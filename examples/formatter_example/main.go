package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := formatterExample(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func formatterExample() (err error) {

	// Create a logger
	log.AddLogger("std", log.NewLogger())

	// Write to all loggers
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	// Change the logger's formatter
	log.GetLogger("std").Options(
		log.WithTimestampFormat(time.RFC1123),
		log.WithFormatter(func(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {
			return fmt.Sprintf("%s - %s - {%s} %s", timestamp, logLevel, tags, fmt.Sprintf("%v", variables)), nil
		}),
	)

	// Write to all loggers again
	log.Fatal(errors.New("Fatal log"))
	log.Error(errors.New("Error log"))
	log.Warn("Warn log")
	log.Info("Info log")
	log.Debug("Debug log", []string{"Random", "string", "array"})
	log.Trace("Trace log", true, 42, 3.14159)

	return
}
