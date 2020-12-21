package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	log "gopkg.in/pd93/plog.v0"
)

func main() {
	if err := FormatterExample(); err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}
}

// FormatterExample will print to stdout using a custom log format.
// If PLogs default formatters aren't quite right for your project, you can provide your own.
// You can do this during logger creation (`log.NewLogger()`) or using the `logger.Options()` method.
// Either way, you do this by providing the `plog.WithFormatter()` functional option as an argument.
// All formatters must take a timestamp, log level, a list of variables and a list of tags as arguments.
// It is up to you to decide how or if you display this data.
func FormatterExample() (err error) {

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
