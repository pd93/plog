package plog

var loggers = make(map[string]*Logger)

//
// Loggers
//

// AddStdLogger creates and adds a new instance of logger with the default config
// See `func NewStdConfig` for more details
func AddStdLogger(name string) {
	loggers[name] = &Logger{
		name:   name,
		config: NewStdConfig(),
	}
}

// AddLogger creates and adds a new instance of Logger with the provided config
// See `type Config` for more details
func AddLogger(name string, config *Config) {
	loggers[name] = &Logger{
		name:   name,
		config: config,
	}
}

//
// Fatal logging (Level 1)
//

// Fatal will print an error message
func Fatal(err error) {
	NewLogf(FatalLevel, "%v", err).WriteToAll()
}

//
// Error logging (Level 2)
//

// Error will print an error message
func Error(err error) {
	NewLogf(ErrorLevel, "%v", err).WriteToAll()
}

//
// Warn logging (Level 3)
//

// Warn will print an object at warn level
func Warn(params ...interface{}) {
	NewLog(WarnLevel, params...).WriteToAll()
}

// Warnf will format and print a message at warn level
func Warnf(message string, params ...interface{}) {
	NewLogf(WarnLevel, message, params...).WriteToAll()
}

//
// Info logging (Level 4)
//

// Info will print an object at info level
func Info(params ...interface{}) {
	NewLog(InfoLevel, params...).WriteToAll()
}

// Infof will format and print a message at info level
func Infof(message string, params ...interface{}) {
	NewLogf(InfoLevel, message, params...).WriteToAll()
}

//
// Debug logging (Level 5)
//

// Debug will print an object at debug level
func Debug(params ...interface{}) {
	NewLog(DebugLevel, params...).WriteToAll()
}

// Debugf will format and print a message at debug level
func Debugf(message string, params ...interface{}) {
	NewLogf(DebugLevel, message, params...).WriteToAll()
}

//
// Trace logging (Level 6)
//

// Trace will print an object at debug level
func Trace(params ...interface{}) {
	NewLog(TraceLevel, params...).WriteToAll()
}

// Tracef will format and print a message at debug level
func Tracef(message string, params ...interface{}) {
	NewLogf(TraceLevel, message, params...).WriteToAll()
}
