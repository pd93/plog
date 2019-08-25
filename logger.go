package plog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// A Logger is a channel for writing logs
type Logger struct {
	Output          io.Writer
	LogLevel        LogLevel
	LogFormat       LogFormat
	TimestampFormat string
	ColorLogging    bool
}

// NewStdLogger creates and returns an instance of Logger with the default settings
// This logger will output to stdout in text format with colors enabled and the RFC3339 timestamp format
func NewStdLogger() (logger *Logger) {
	return &Logger{
		Output:          os.Stdout,
		LogLevel:        InfoLevel,
		LogFormat:       TextFormat,
		TimestampFormat: time.RFC3339,
		ColorLogging:    true,
	}
}

// Validate will set default values for uninitialised values
// It also check whether or not the logger is configured correctly and return any errors it finds
func (logger *Logger) Validate() (err error) {

	// Set the default out to stdout
	if logger.Output == nil {
		logger.Output = os.Stdout
	}

	// Set the default timestamp format to RFC3339
	if logger.TimestampFormat == "" {
		logger.TimestampFormat = time.RFC3339
	}

	// Turn on coloured logs for the text format
	if logger.LogFormat == TextFormat {
		logger.ColorLogging = true
	}

	// Check if the log level is valid
	if logger.LogLevel.String() == "" {
		return errors.New("Invalid log level")
	}

	// Check if the log format is valid
	if logger.LogFormat.String() == "" {
		return errors.New("Invalid log format")
	}

	return
}

// SetOutput allows you to change where the logs are being output to.
// Examples include 'os.File', 'os.Stdout', 'os.Stderr', 'os.Stdin' or any other writer.
func (logger *Logger) SetOutput(output io.Writer) {
	logger.Output = output
}

// SetLogLevel will set the level of the log message
func (logger *Logger) SetLogLevel(logLevel LogLevel) {
	logger.LogLevel = logLevel
}

// SetLogFormat will set the format of the log message
func (logger *Logger) SetLogFormat(logFormat LogFormat) {
	logger.LogFormat = logFormat
}

// SetTimestampFormat allows the user to specify a custom timestamp format.
// The default format is 'time.RFC3339'.
// You can find the documentation on time formatting in Golang here: https://golang.org/pkg/time/#Time.Format
func (logger *Logger) SetTimestampFormat(timestampFormat string) {
	logger.TimestampFormat = timestampFormat
}

// SetColorLogging will enable/disable colored logging.
func (logger *Logger) SetColorLogging(colorLogging bool) {
	logger.ColorLogging = colorLogging
}

// Write will add a log message to the logger
func (logger *Logger) Write(log *Log) (err error) {

	var message, body string

	// Check if we need to log this message or not
	if logger.LogLevel >= log.LogLevel {

		switch logger.LogFormat {

		//
		// Text logger
		//

		case TextFormat:

			// Render the timestamp and log level strings
			timestamp := log.Timestamp.Format(logger.TimestampFormat)
			logLevel := strings.ToUpper(log.LogLevel.String())

			// Check if colored logging is enabled
			if logger.ColorLogging {

				// Set the color of the text
				switch log.LogLevel {
				case ErrorLevel, FatalLevel:
					logLevel = color(logLevel, FgRed)
				case WarnLevel:
					logLevel = color(logLevel, FgYellow)
				case InfoLevel:
					logLevel = color(logLevel, FgGreen)
				case DebugLevel:
					logLevel = color(logLevel, FgCyan)
				case TraceLevel:
					logLevel = color(logLevel, FgBlue)
				}
			}

			// Stringify the body
			for _, item := range log.Body {
				body += fmt.Sprintf(" %v", item)
			}

			// Format the message
			message = fmt.Sprintf("%s [%s] %s", timestamp, logLevel, body)

		//
		// JSON logger
		//

		case JSONFormat:
			var b []byte
			b, err = json.Marshal(log)
			message = string(b)
		}

		// Print the message to the output writer
		fmt.Fprintln(logger.Output, message)
	}

	return
}
