package plog

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

	loggers[name] = logger
}

// DeleteLogger removes the specified logger from PLog
func DeleteLogger(name string) {
	delete(loggers, name)
}

// GetLogger returns the specified logger
func GetLogger(name string) *Logger {
	return loggers[name]
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
func Warn(params ...interface{}) {
	loggers.Write(NewLog(WarnLevel, params...))
}

// Warnf will print a formatted message to all loggers at warn level
func Warnf(message string, params ...interface{}) {
	loggers.Write(NewLogf(WarnLevel, message, params...))
}

//
// Info logging (Level 4)
//

// Info will print any number of variables to all loggers at info level
func Info(params ...interface{}) {
	loggers.Write(NewLog(InfoLevel, params...))
}

// Infof will print a formatted message to all loggers at info level
func Infof(message string, params ...interface{}) {
	loggers.Write(NewLogf(InfoLevel, message, params...))
}

//
// Debug logging (Level 5)
//

// Debug will print any number of variables to all loggers at debug level
func Debug(params ...interface{}) {
	loggers.Write(NewLog(DebugLevel, params...))
}

// Debugf will print a formatted message to all loggers at debug level
func Debugf(message string, params ...interface{}) {
	loggers.Write(NewLogf(DebugLevel, message, params...))
}

//
// Trace logging (Level 6)
//

// Trace will print any number of variables to all loggers at debug level
func Trace(params ...interface{}) {
	loggers.Write(NewLog(TraceLevel, params...))
}

// Tracef will print a formatted message to all loggers at debug level
func Tracef(message string, params ...interface{}) {
	loggers.Write(NewLogf(TraceLevel, message, params...))
}
