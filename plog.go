package plog

var loggers = make(map[string]*Logger)

//
// Loggers
//

// AddLogger adds the provided logger to PLog
// See `type Logger` for more details
func AddLogger(name string, logger *Logger) {
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
	NewLogf(FatalLevel, "%v", err).Write()
}

//
// Error logging (Level 2)
//

// Error will print an error message
func Error(err error) {
	NewLogf(ErrorLevel, "%v", err).Write()
}

//
// Warn logging (Level 3)
//

// Warn will print an object at warn level
func Warn(params ...interface{}) {
	NewLog(WarnLevel, params...).Write()
}

// Warnf will format and print a message at warn level
func Warnf(message string, params ...interface{}) {
	NewLogf(WarnLevel, message, params...).Write()
}

//
// Info logging (Level 4)
//

// Info will print an object at info level
func Info(params ...interface{}) {
	NewLog(InfoLevel, params...).Write()
}

// Infof will format and print a message at info level
func Infof(message string, params ...interface{}) {
	NewLogf(InfoLevel, message, params...).Write()
}

//
// Debug logging (Level 5)
//

// Debug will print an object at debug level
func Debug(params ...interface{}) {
	NewLog(DebugLevel, params...).Write()
}

// Debugf will format and print a message at debug level
func Debugf(message string, params ...interface{}) {
	NewLogf(DebugLevel, message, params...).Write()
}

//
// Trace logging (Level 6)
//

// Trace will print an object at debug level
func Trace(params ...interface{}) {
	NewLog(TraceLevel, params...).Write()
}

// Tracef will format and print a message at debug level
func Tracef(message string, params ...interface{}) {
	NewLogf(TraceLevel, message, params...).Write()
}
