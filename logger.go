package plog

import (
	"fmt"
	"io"
	"strings"
)

// A Logger is a channel for writing logs
type Logger struct {
	output          io.Writer
	logLevel        LogLevel
	logFormat       LogFormat
	timestampFormat string
	colorLogging    bool
}

//
// Getters
//

// Output will return the logger's current output
func (logger *Logger) Output() io.Writer {
	return logger.output
}

// LogLevel will return the logger's current log level
func (logger *Logger) LogLevel() LogLevel {
	return logger.logLevel
}

// LogFormat will return the logger's current log format
func (logger *Logger) LogFormat() LogFormat {
	return logger.logFormat
}

// TimestampFormat will return the logger's current timestamp format
func (logger *Logger) TimestampFormat() string {
	return logger.timestampFormat
}

// ColorLogging will return whether or not color logging is enabled
func (logger *Logger) ColorLogging() bool {
	return logger.colorLogging
}

//
// Setters
//

// SetOutput allows you to change where the logs are being output to.
// Examples include 'os.File', 'os.Stdout', 'bytes.Buffer' or any other writer.
func (logger *Logger) SetOutput(output io.Writer) {
	logger.output = output
}

// SetLogLevel will set the level of the log message
func (logger *Logger) SetLogLevel(logLevel LogLevel) {
	logger.logLevel = logLevel
}

// SetLogFormat will set the format of the log message
func (logger *Logger) SetLogFormat(logFormat LogFormat) {
	logger.logFormat = logFormat
}

// SetTimestampFormat allows the user to specify a custom timestamp format.
// The default format is 'time.RFC3339'.
// You can find the documentation on time formatting in Golang here: https://golang.org/pkg/time/#Time.Format
func (logger *Logger) SetTimestampFormat(timestampFormat string) {
	logger.timestampFormat = timestampFormat
}

// SetColorLogging will enable/disable colored logging.
func (logger *Logger) SetColorLogging(colorLogging bool) {
	logger.colorLogging = colorLogging
}

//
// Fatal logging (Level 1)
//

// Fatal will print a fatal error message
func (logger *Logger) Fatal(err error) {
	logger.write(newLogf(FatalLevel, "%v", err))
}

//
// Error logging (Level 2)
//

// Error will print a non-fatal error message
func (logger *Logger) Error(err error) {
	logger.write(newLogf(ErrorLevel, "%v", err))
}

//
// Warn logging (Level 3)
//

// Warn will print any number of variables at warn level
func (logger *Logger) Warn(variables ...interface{}) {
	logger.write(newLog(WarnLevel, variables...))
}

// Warnf will print a formatted message at warn level
func (logger *Logger) Warnf(format string, variables ...interface{}) {
	logger.write(newLogf(WarnLevel, format, variables...))
}

//
// Info logging (Level 4)
//

// Info will print any number of variables at info level
func (logger *Logger) Info(variables ...interface{}) {
	logger.write(newLog(InfoLevel, variables...))
}

// Infof will print a formatted message at info level
func (logger *Logger) Infof(format string, variables ...interface{}) {
	logger.write(newLogf(InfoLevel, format, variables...))
}

//
// Debug logging (Level 5)
//

// Debug will print any number of variables at debug level
func (logger *Logger) Debug(variables ...interface{}) {
	logger.write(newLog(DebugLevel, variables...))
}

// Debugf will print a formatted message at debug level
func (logger *Logger) Debugf(format string, variables ...interface{}) {
	logger.write(newLogf(DebugLevel, format, variables...))
}

//
// Trace logging (Level 6)
//

// Trace will print any number of variables at debug level
func (logger *Logger) Trace(variables ...interface{}) {
	logger.write(newLog(TraceLevel, variables...))
}

// Tracef will print a formatted message at debug level
func (logger *Logger) Tracef(format string, variables ...interface{}) {
	logger.write(newLogf(TraceLevel, format, variables...))
}

//
// Writer
//

// write will add a log message to the logger
func (logger *Logger) write(l *log) {

	// Check if we need to log this message or not
	if logger.logLevel >= l.logLevel {

		// Render the timestamp and log level strings
		timestamp := l.timestamp.Format(logger.timestampFormat)
		logLevel := strings.ToUpper(l.logLevel.String())

		// Check if colored logging is enabled
		if logger.colorLogging {

			// Set the color of the text
			switch l.logLevel {
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
		variables := make([]string, len(l.variables))
		for i, variable := range l.variables {
			variables[i] = fmt.Sprintf("%v", variable)
		}
		message := strings.Join(variables, " ")

		// Create the output string
		var outputString string

		switch logger.logFormat {

		case TextFormat:
			outputString = fmt.Sprintf("%s [%s] %s", timestamp, logLevel, message)

		case JSONFormat:
			outputString = fmt.Sprintf(`{ "timestamp": "%s", "logLevel": "%s", "message": "%s" }`, timestamp, logLevel, message)

		case CSVFormat:
			outputString = fmt.Sprintf(`%s,%s,%s`, timestamp, logLevel, message)
		}

		// Print the message to the output writer
		fmt.Fprintln(logger.output, outputString)
	}

	return
}
