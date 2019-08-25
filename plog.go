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

// Fatal will print an error message
func Fatal(err error) {
	loggers.Write(NewLogf(FatalLevel, "%v", err))
}

//
// Error logging (Level 2)
//

// Error will print an error message
func Error(err error) {
	loggers.Write(NewLogf(ErrorLevel, "%v", err))
}

//
// Warn logging (Level 3)
//

// Warn will print an object at warn level
func Warn(params ...interface{}) {
	loggers.Write(NewLog(WarnLevel, params...))
}

// Warnf will format and print a message at warn level
func Warnf(message string, params ...interface{}) {
	loggers.Write(NewLogf(WarnLevel, message, params...))
}

//
// Info logging (Level 4)
//

// Info will print an object at info level
func Info(params ...interface{}) {
	loggers.Write(NewLog(InfoLevel, params...))
}

// Infof will format and print a message at info level
func Infof(message string, params ...interface{}) {
	loggers.Write(NewLogf(InfoLevel, message, params...))
}

//
// Debug logging (Level 5)
//

// Debug will print an object at debug level
func Debug(params ...interface{}) {
	loggers.Write(NewLog(DebugLevel, params...))
}

// Debugf will format and print a message at debug level
func Debugf(message string, params ...interface{}) {
	loggers.Write(NewLogf(DebugLevel, message, params...))
}

//
// Trace logging (Level 6)
//

// Trace will print an object at debug level
func Trace(params ...interface{}) {
	loggers.Write(NewLog(TraceLevel, params...))
}

// Tracef will format and print a message at debug level
func Tracef(message string, params ...interface{}) {
	loggers.Write(NewLogf(TraceLevel, message, params...))
}
