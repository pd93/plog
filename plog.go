package plog

import (
	"fmt"
	"os"
)

var loggers = make(loggerMap)

//
// Loggers
//

// NewLogger creates and returns an instance of Logger with the default variables
func NewLogger() *Logger {
	return &Logger{}
}

// NewJSONFileLogger creates and returns an instance of Logger which will write to the specified file
// The log level is set to TraceLevel (log everything) and color logging is disabled
// The logs will be written in JSON format, but the file does not need to end in '.json'
func NewJSONFileLogger(file *os.File) *Logger {
	return &Logger{
		output:    file,
		logLevel:  TraceLevel,
		logFormat: JSONFormat,
	}
}

// NewCSVFileLogger creates and returns an instance of Logger which will write to the specified file
// The log level is set to TraceLevel (log everything) and color logging is disabled
// The logs will be written in CSV format, but the file does not need to end in '.csv'
func NewCSVFileLogger(file *os.File) *Logger {
	return &Logger{
		output:    file,
		logLevel:  TraceLevel,
		logFormat: CSVFormat,
	}
}

// AddLogger adds the provided logger to PLog
// See `type Logger` for more details
func AddLogger(name string, logger *Logger) {

	// Check that the logger is valid are set the default values
	if err := logger.validate(); err != nil {
		panic(err)
	}

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
