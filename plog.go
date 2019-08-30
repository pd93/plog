package plog

import (
	"fmt"
	"io"
	"os"
	"time"
)

//
// Loggers
//

// Global loggers variable
var loggers = make(loggerMap)

// NewLogger creates and returns an instance of Logger with the default variables
func NewLogger() *Logger {
	return &Logger{
		output:          os.Stdout,
		logLevel:        InfoLevel,
		logFormat:       TextFormat,
		timestampFormat: time.RFC3339,
		colorLogging:    true,
	}
}

// NewJSONFileLogger creates and returns an instance of Logger which will write to the specified file
// The log level is set to TraceLevel (log everything) and color logging is disabled
// The logs will be written in JSON format, but the file does not need to end in '.json'
func NewJSONFileLogger(output io.Writer) *Logger {
	return &Logger{
		output:          output,
		logLevel:        TraceLevel,
		logFormat:       JSONFormat,
		timestampFormat: time.RFC3339,
		colorLogging:    false,
	}
}

// NewCSVFileLogger creates and returns an instance of Logger which will write to the specified file
// The log level is set to TraceLevel (log everything) and color logging is disabled
// The logs will be written in CSV format, but the file does not need to end in '.csv'
func NewCSVFileLogger(output io.Writer) *Logger {
	return &Logger{
		output:          output,
		logLevel:        TraceLevel,
		logFormat:       CSVFormat,
		timestampFormat: time.RFC3339,
		colorLogging:    false,
	}
}

// AddLogger adds the provided logger to PLog
// See `type Logger` for more details
func AddLogger(name string, logger *Logger) {

	// Check if the logger name is already used
	if _, exists := loggers[name]; exists {
		panic(fmt.Errorf("Logger with the name: '%s' already exists", name))
	}

	loggers[name] = logger
}

// GetLogger returns the specified logger
func GetLogger(name string) *Logger {

	// Check if the logger exists
	if val, exists := loggers[name]; !exists || val == nil {
		panic(fmt.Errorf("Cannot return non-existent logger: '%s'", name))
	}

	return loggers[name]
}

// DeleteLogger removes the specified logger from PLog
func DeleteLogger(name string) {

	// Check if the logger exists
	if val, exists := loggers[name]; !exists || val == nil {
		panic(fmt.Errorf("Cannot delete non-existent logger: '%s'", name))
	}

	delete(loggers, name)
}

//
// Fatal logging (Level 1)
//

// Fatal will print a fatal error message to all loggers
func Fatal(err error) {
	loggers.write(newLogf(FatalLevel, "%v", err))
}

//
// Error logging (Level 2)
//

// Error will print a non-fatal error message to all loggers
func Error(err error) {
	loggers.write(newLogf(ErrorLevel, "%v", err))
}

//
// Warn logging (Level 3)
//

// Warn will print any number of variables to all loggers at warn level
func Warn(variables ...interface{}) {
	loggers.write(newLog(WarnLevel, variables...))
}

// Warnf will print a formatted message to all loggers at warn level
func Warnf(format string, variables ...interface{}) {
	loggers.write(newLogf(WarnLevel, format, variables...))
}

//
// Info logging (Level 4)
//

// Info will print any number of variables to all loggers at info level
func Info(variables ...interface{}) {
	loggers.write(newLog(InfoLevel, variables...))
}

// Infof will print a formatted message to all loggers at info level
func Infof(format string, variables ...interface{}) {
	loggers.write(newLogf(InfoLevel, format, variables...))
}

//
// Debug logging (Level 5)
//

// Debug will print any number of variables to all loggers at debug level
func Debug(variables ...interface{}) {
	loggers.write(newLog(DebugLevel, variables...))
}

// Debugf will print a formatted message to all loggers at debug level
func Debugf(format string, variables ...interface{}) {
	loggers.write(newLogf(DebugLevel, format, variables...))
}

//
// Trace logging (Level 6)
//

// Trace will print any number of variables to all loggers at debug level
func Trace(variables ...interface{}) {
	loggers.write(newLog(TraceLevel, variables...))
}

// Tracef will print a formatted message to all loggers at debug level
func Tracef(format string, variables ...interface{}) {
	loggers.write(newLogf(TraceLevel, format, variables...))
}

//
// Convenience functions
//

// SetOutput will loop through all the loggers and set where the logs are being output to.
// Examples include 'os.File', 'os.Stdout', 'bytes.Buffer' or any other writer.
func SetOutput(output io.Writer) {
	for _, logger := range loggers {
		logger.SetOutput(output)
	}
}

// SetLogLevel will loop through all the loggers and set the level of the log message
func SetLogLevel(logLevel LogLevel) {
	for _, logger := range loggers {
		logger.SetLogLevel(logLevel)
	}
}

// SetLogFormat will loop through all the loggers and set the format of the log message
func SetLogFormat(logFormat LogFormat) {
	for _, logger := range loggers {
		logger.SetLogFormat(logFormat)
	}
}

// SetTimestampFormat will loop through all the loggers and set the timestampFormat setting
// The default format is 'time.RFC3339'.
// You can find the documentation on time formatting in Golang here: https://golang.org/pkg/time/#Time.Format
func SetTimestampFormat(timestampFormat string) {
	for _, logger := range loggers {
		logger.SetTimestampFormat(timestampFormat)
	}
}

// SetColorLogging will loop through all the loggers and enable/disable colored logging.
func SetColorLogging(colorLogging bool) {
	for _, logger := range loggers {
		logger.SetColorLogging(colorLogging)
	}
}
