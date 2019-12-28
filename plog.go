package plog

import (
	"fmt"
)

//
// Loggers
//

// Global loggers variable
var loggers = make(loggerMap)

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

// Options will apply the given options to all loggers
// Any number of functional options can be passed to this method
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options
func Options(opts ...LoggerOption) {
	for _, logger := range loggers {
		logger.Options(opts...)
	}
}

//
// Fatal logging (Level 1)
//

// Fatal will print a fatal error message to all loggers
func Fatal(err error) {
	loggers.write(newLogf(FatalLevel, "%v", err))
}

// Fatalf will print a formatted, non-fatal error message
func Fatalf(format string, err error) {
	loggers.write(newLogf(FatalLevel, format, err))
}

// TFatal will print a fatal error message and meta-tag the log
func TFatal(tags Tags, err error) {
	loggers.write(newTLogf(FatalLevel, tags, "%v", err))
}

// TFatalf will print a formatted, fatal error message and meta-tag the log
func TFatalf(tags Tags, format string, err error) {
	loggers.write(newTLogf(FatalLevel, tags, format, err))
}

//
// Error logging (Level 2)
//

// Error will print a non-fatal error message to all loggers
func Error(err error) {
	loggers.write(newLogf(ErrorLevel, "%v", err))
}

// Errorf will print a formatted, non-fatal error message
func Errorf(format string, err error) {
	loggers.write(newLogf(ErrorLevel, format, err))
}

// TError will print a non-fatal error message and meta-tag the log
func TError(tags Tags, err error) {
	loggers.write(newTLogf(ErrorLevel, tags, "%v", err))
}

// TErrorf will print a formatted, non-fatal error message and meta-tag the log
func TErrorf(tags Tags, format string, err error) {
	loggers.write(newTLogf(ErrorLevel, tags, format, err))
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

// TWarn will print any number of variables at warn level and meta-tag the log
func TWarn(tags Tags, variables ...interface{}) {
	loggers.write(newTLog(WarnLevel, tags, variables...))
}

// TWarnf will print a formatted message at warn level and meta-tag the log
func TWarnf(tags Tags, format string, variables ...interface{}) {
	loggers.write(newTLogf(WarnLevel, tags, format, variables...))
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

// TInfo will print any number of variables at info level and meta-tag the log
func TInfo(tags Tags, variables ...interface{}) {
	loggers.write(newTLog(InfoLevel, tags, variables...))
}

// TInfof will print a formatted message at info level and meta-tag the log
func TInfof(tags Tags, format string, variables ...interface{}) {
	loggers.write(newTLogf(InfoLevel, tags, format, variables...))
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

// TDebug will print any number of variables at debug level and meta-tag the log
func TDebug(tags Tags, variables ...interface{}) {
	loggers.write(newTLog(DebugLevel, tags, variables...))
}

// TDebugf will print a formatted message at debug level and meta-tag the log
func TDebugf(tags Tags, format string, variables ...interface{}) {
	loggers.write(newTLogf(DebugLevel, tags, format, variables...))
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

// TTrace will print any number of variables at trace level and meta-tag the log
func TTrace(tags Tags, variables ...interface{}) {
	loggers.write(newTLog(TraceLevel, tags, variables...))
}

// TTracef will print a formatted message at trace level and meta-tag the log
func TTracef(tags Tags, format string, variables ...interface{}) {
	loggers.write(newTLogf(TraceLevel, tags, format, variables...))
}
