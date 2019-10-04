package plog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// A Logger is a channel for writing logs
type Logger struct {
	output           io.Writer
	logLevel         LogLevel
	logFormat        LogFormat
	timestampFormat  string
	colorLogging     bool
	logLevelColorMap LogLevelColorMap
	tagColorMap      TagColorMap
}

//
// Constructors
//

// NewLogger creates and returns an instance of Logger with the default variables
func NewLogger() *Logger {
	return &Logger{
		output:           os.Stdout,
		logLevel:         InfoLevel,
		logFormat:        TextFormat,
		timestampFormat:  time.RFC3339,
		colorLogging:     true,
		logLevelColorMap: NewLogLevelColorMap(),
		tagColorMap:      NewTagColorMap(),
	}
}

// NewJSONFileLogger creates and returns an instance of Logger which will write to the specified file
// The log level is set to TraceLevel (log everything) and color logging is disabled
// The logs will be written in JSON format, but the file does not need to end in '.json'
func NewJSONFileLogger(output io.Writer) *Logger {
	return &Logger{
		output:           output,
		logLevel:         TraceLevel,
		logFormat:        JSONFormat,
		timestampFormat:  time.RFC3339,
		colorLogging:     false,
		logLevelColorMap: NewLogLevelColorMap(),
		tagColorMap:      NewTagColorMap(),
	}
}

// NewCSVFileLogger creates and returns an instance of Logger which will write to the specified file
// The log level is set to TraceLevel (log everything) and color logging is disabled
// The logs will be written in CSV format, but the file does not need to end in '.csv'
func NewCSVFileLogger(output io.Writer) *Logger {
	return &Logger{
		output:           output,
		logLevel:         TraceLevel,
		logFormat:        CSVFormat,
		timestampFormat:  time.RFC3339,
		colorLogging:     false,
		logLevelColorMap: NewLogLevelColorMap(),
		tagColorMap:      NewTagColorMap(),
	}
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

// LogLevelColorMap will return the logger's text attributes for each log level
func (logger *Logger) LogLevelColorMap() LogLevelColorMap {
	return logger.logLevelColorMap
}

// TagColorMap will return the logger's text attributes for each tag
func (logger *Logger) TagColorMap() TagColorMap {
	return logger.tagColorMap
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

// SetLogLevelColorMap set the colors for each log level.
func (logger *Logger) SetLogLevelColorMap(logLevelColorMap LogLevelColorMap) {
	logger.logLevelColorMap = logLevelColorMap
}

// SetTagColorMap set the colors for each tag.
func (logger *Logger) SetTagColorMap(tagColorMap TagColorMap) {
	logger.tagColorMap = tagColorMap
}

//
// Fatal logging (Level 1)
//

// Fatal will print a fatal error message
func (logger *Logger) Fatal(err error) {
	logger.write(newLogf(FatalLevel, "%v", err))
}

// Fatalf will print a formatted, non-fatal error message
func (logger *Logger) Fatalf(format string, err error) {
	logger.write(newLogf(FatalLevel, format, err))
}

// TFatal will print a fatal error message and meta-tag the log
func (logger *Logger) TFatal(tags Tags, err error) {
	logger.write(newTLogf(FatalLevel, tags, "%v", err))
}

// TFatalf will print a formatted, fatal error message and meta-tag the log
func (logger *Logger) TFatalf(tags Tags, format string, err error) {
	logger.write(newTLogf(FatalLevel, tags, format, err))
}

//
// Error logging (Level 2)
//

// Error will print a non-fatal error message
func (logger *Logger) Error(err error) {
	logger.write(newLogf(ErrorLevel, "%v", err))
}

// Errorf will print a formatted, non-fatal error message
func (logger *Logger) Errorf(format string, err error) {
	logger.write(newLogf(ErrorLevel, format, err))
}

// TError will print a non-fatal error message and meta-tag the log
func (logger *Logger) TError(tags Tags, err error) {
	logger.write(newTLogf(ErrorLevel, tags, "%v", err))
}

// TErrorf will print a formatted, non-fatal error message and meta-tag the log
func (logger *Logger) TErrorf(tags Tags, format string, err error) {
	logger.write(newTLogf(ErrorLevel, tags, format, err))
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

// TWarn will print any number of variables at warn level and meta-tag the log
func (logger *Logger) TWarn(tags Tags, variables ...interface{}) {
	logger.write(newTLog(WarnLevel, tags, variables...))
}

// TWarnf will print a formatted message at warn level and meta-tag the log
func (logger *Logger) TWarnf(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(WarnLevel, tags, format, variables...))
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

// TInfo will print any number of variables at info level and meta-tag the log
func (logger *Logger) TInfo(tags Tags, variables ...interface{}) {
	logger.write(newTLog(InfoLevel, tags, variables...))
}

// TInfof will print a formatted message at info level and meta-tag the log
func (logger *Logger) TInfof(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(InfoLevel, tags, format, variables...))
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

// TDebug will print any number of variables at debug level and meta-tag the log
func (logger *Logger) TDebug(tags Tags, variables ...interface{}) {
	logger.write(newTLog(DebugLevel, tags, variables...))
}

// TDebugf will print a formatted message at debug level and meta-tag the log
func (logger *Logger) TDebugf(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(DebugLevel, tags, format, variables...))
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

// TTrace will print any number of variables at trace level and meta-tag the log
func (logger *Logger) TTrace(tags Tags, variables ...interface{}) {
	logger.write(newTLog(TraceLevel, tags, variables...))
}

// TTracef will print a formatted message at trace level and meta-tag the log
func (logger *Logger) TTracef(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(TraceLevel, tags, format, variables...))
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
		logLevel := l.logLevel.text(logger.colorLogging, logger.logLevelColorMap)

		// Stringify the variables
		variables := make([]string, len(l.variables))
		for i, variable := range l.variables {
			variables[i] = fmt.Sprintf("%v", variable)
		}
		message := strings.Join(variables, " ")

		var tags string
		var outputString string

		// Create the output string
		switch logger.logFormat {

		case TextFormat:
			tags = l.tags.text(logger.colorLogging, logger.tagColorMap)
			outputString = fmt.Sprintf("%s [%s] %s %s", timestamp, logLevel, message, tags)

		case JSONFormat:
			tags = l.tags.json(logger.colorLogging, logger.tagColorMap)
			outputString = fmt.Sprintf(`{ "timestamp": "%s", "logLevel": "%s", "message": "%s", "tags": %s }`, timestamp, logLevel, message, tags)

		case CSVFormat:
			tags = l.tags.csv(logger.colorLogging, logger.tagColorMap)
			outputString = fmt.Sprintf(`%s,%s,%s,%s`, timestamp, logLevel, message, tags)
		}

		// Print the message to the output writer
		fmt.Fprintln(logger.output, outputString)
	}

	return
}
