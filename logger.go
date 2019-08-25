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

//
// Setters
//

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

//
// Writer
//

// Write will add a log message to the logger
func (logger *Logger) Write(log *Log) (err error) {

	// Check if we need to log this message or not
	if logger.LogLevel >= log.LogLevel {

		var message string

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

			// Stringify the variables
			variables := make([]string, len(log.Body))
			for i, variable := range log.Body {
				variables[i] = fmt.Sprintf("%v", variable)
			}

			// Format the message
			message = fmt.Sprintf("%s [%s] %s", timestamp, logLevel, strings.Join(variables, " "))

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

//
// Fatal logging (Level 1)
//

// Fatal will print a fatal error message
func (logger *Logger) Fatal(err error) {
	logger.Write(NewLogf(FatalLevel, "%v", err))
}

//
// Error logging (Level 2)
//

// Error will print a non-fatal error message
func (logger *Logger) Error(err error) {
	logger.Write(NewLogf(ErrorLevel, "%v", err))
}

//
// Warn logging (Level 3)
//

// Warn will print any number of variables at warn level
func (logger *Logger) Warn(params ...interface{}) {
	logger.Write(NewLog(WarnLevel, params...))
}

// Warnf will print a formatted message at warn level
func (logger *Logger) Warnf(message string, params ...interface{}) {
	logger.Write(NewLogf(WarnLevel, message, params...))
}

//
// Info logging (Level 4)
//

// Info will print any number of variables at info level
func (logger *Logger) Info(params ...interface{}) {
	logger.Write(NewLog(InfoLevel, params...))
}

// Infof will print a formatted message at info level
func (logger *Logger) Infof(message string, params ...interface{}) {
	logger.Write(NewLogf(InfoLevel, message, params...))
}

//
// Debug logging (Level 5)
//

// Debug will print any number of variables at debug level
func (logger *Logger) Debug(params ...interface{}) {
	logger.Write(NewLog(DebugLevel, params...))
}

// Debugf will print a formatted message at debug level
func (logger *Logger) Debugf(message string, params ...interface{}) {
	logger.Write(NewLogf(DebugLevel, message, params...))
}

//
// Trace logging (Level 6)
//

// Trace will print any number of variables at debug level
func (logger *Logger) Trace(params ...interface{}) {
	logger.Write(NewLog(TraceLevel, params...))
}

// Tracef will print a formatted message at debug level
func (logger *Logger) Tracef(message string, params ...interface{}) {
	logger.Write(NewLogf(TraceLevel, message, params...))
}
