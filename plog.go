package plog

import "fmt"

var loggers Loggers = make(map[string]*Logger)

//
// Loggers
//

// AddLogger adds the provided logger to PLog
// See `type Logger` for more details
func AddLogger(name string, logger *Logger) {

	// Check that the logger is valid are set the default values
	if err := logger.Validate(); err != nil {
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
	loggers.Write(NewLogf(FatalLevel, "%v", err))
}

//
// Error logging (Level 2)
//

// Error will print a non-fatal error message to all loggers
func Error(err error) {
	loggers.Write(NewLogf(ErrorLevel, "%v", err))
}

//
// Warn logging (Level 3)
//

// Warn will print any number of variables to all loggers at warn level
func Warn(variables ...interface{}) {
	loggers.Write(NewLog(WarnLevel, variables...))
}

// Warnf will print a formatted message to all loggers at warn level
func Warnf(format string, variables ...interface{}) {
	loggers.Write(NewLogf(WarnLevel, format, variables...))
}

//
// Info logging (Level 4)
//

// Info will print any number of variables to all loggers at info level
func Info(variables ...interface{}) {
	loggers.Write(NewLog(InfoLevel, variables...))
}

// Infof will print a formatted message to all loggers at info level
func Infof(format string, variables ...interface{}) {
	loggers.Write(NewLogf(InfoLevel, format, variables...))
}

//
// Debug logging (Level 5)
//

// Debug will print any number of variables to all loggers at debug level
func Debug(variables ...interface{}) {
	loggers.Write(NewLog(DebugLevel, variables...))
}

// Debugf will print a formatted message to all loggers at debug level
func Debugf(format string, variables ...interface{}) {
	loggers.Write(NewLogf(DebugLevel, format, variables...))
}

//
// Trace logging (Level 6)
//

// Trace will print any number of variables to all loggers at debug level
func Trace(variables ...interface{}) {
	loggers.Write(NewLog(TraceLevel, variables...))
}

// Tracef will print a formatted message to all loggers at debug level
func Tracef(format string, variables ...interface{}) {
	loggers.Write(NewLogf(TraceLevel, format, variables...))
}
